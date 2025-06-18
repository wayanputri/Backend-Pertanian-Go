package main

import (
	"backend/app/config"
	"backend/app/database"
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

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	ctx := context.Background()
	db := database.ProstgresConfig()

	dbGorm := database.InitPostgresGorm()

	dbr := database.InitRedis(ctx)
	clientM, collectionM := database.InitMongo(ctx)
	providerNew := repositori.New(db, dbr, collectionM, dbGorm)

	defer db.Close()
	defer clientM.Disconnect(ctx)
	serverNew := handler.New(providerNew)
	// Serve Swagger UI from "swagger/swagger-ui" folder
	e.Static("/api/docs", "swagger/swagger-ui")

	// Serve Swagger JSON from "swagger/api.swagger.json"
	e.GET("/api/api.swagger.json", func(c echo.Context) error {
		return c.File("swagger/api.swagger.json")
	})

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

	// Integrate gRPC-Gateway with Echo
	e.Any("/api/*", echo.WrapHandler(gwMux))

	// Start HTTP server (for Swagger and gRPC-Gateway)
	log.Println("Starting HTTP server on", cfg.ServerPort)
	e.Logger.Fatal(e.Start(cfg.ServerPort))
}
