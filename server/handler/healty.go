package handler

import (
	"backend/pb"
	"context"
	"fmt"
)

func (s *Server) GetExample(ctx context.Context, req *pb.ExampleRequest2) (*pb.ExampleResponse2, error) {
	s.GetProvider().GetUsers()
	s.GetProvider().SetData(ctx, "coba", "sini dulu")
	resp := s.GetProvider().GetData(ctx, "coba")
	fmt.Println("coba ", resp)

	s.GetProvider().CreateUserMongo(ctx)
	s.GetProvider().GetUserMongo(ctx)
	return &pb.ExampleResponse2{
		Message: fmt.Sprintf("Hello, %s!", req.Name),
	}, nil
}
