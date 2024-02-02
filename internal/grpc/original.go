package grpc

import (
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetOriginalURL -
func (s *Server) GetOriginalURL(_, req *ShortURLRequest) (*OriginalURLResponse, error) {
	short := req.Url

	if len(short) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Empty URL")
	}

	link, err := s.Storage.GetOriginal(short)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid URL")
	}

	return &OriginalURLResponse{
		Original: link.Original,
	}, nil
}
