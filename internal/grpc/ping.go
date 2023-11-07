package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Ping() (*PingResponse, error) {

	err := s.Storage.Ping()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Don't ping Database")
	}

	return &PingResponse{
		Result: "OK",
	}, nil
}
