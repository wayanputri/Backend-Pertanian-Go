package handler

import (
	"backend/pb"
	"backend/server/repositori"
)

// Server struct for gRPC
type Server struct {
	provider *repositori.Provider

	pb.UnimplementedExampleServiceServer
}

func New(db *repositori.Provider) *Server {
	return &Server{
		provider: db,
	}
}

func (s *Server) GetProvider() *repositori.Provider {
	return s.provider
}
