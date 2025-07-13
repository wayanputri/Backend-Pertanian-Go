package main

import (
	"backend/app/config"
	"backend/app/database"
	"backend/lib/external/minio"
	"backend/pb"
	"backend/server/handler"
	"backend/server/repositori"
	"fmt"

	"context"
	"log"
	"net"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc/reflection"
)

func main() {
	// Create Echo instance
	config.LoadConfig()
	cfg := config.AppConfig
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	minio.MinioInit()
	e.POST("/api/upload", minio.UploadFile)

	ctx := context.Background()
	db := database.ProstgresConfig()

	dbGorm := database.InitPostgresGorm()

	dbr := database.InitRedis(ctx)
	clientM, collectionM := database.InitMongo(ctx)
	providerNew := repositori.New(db, dbr, collectionM, dbGorm)

	defer db.Close()
	defer clientM.Disconnect(ctx)
	serverNew := handler.New(providerNew)
	e.Static("/docs/", "swagger/swagger-ui")
	e.File("/swagger.json", "swagger/api.swagger.json")

	// gRPC server
	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterExampleServiceServer(grpcServer, serverNew)

	reflection.Register(grpcServer)
	go func() {
		log.Println("Starting gRPC server...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to start gRPC server: %v", err)
		}
	}()

	// Register gRPC-Gateway handler
	gwMux := runtime.NewServeMux()
	err = pb.RegisterExampleServiceHandlerFromEndpoint(context.Background(), gwMux, fmt.Sprintf("%s%s", cfg.GRPCHost, cfg.GRPCPort), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}

	e.Any("/api/*", echo.WrapHandler(gwMux))

	// Start HTTP server (for Swagger and gRPC-Gateway)
	log.Println("Starting HTTP server on", cfg.ServerPort)
	e.Logger.Fatal(e.Start(cfg.ServerPort))
}
