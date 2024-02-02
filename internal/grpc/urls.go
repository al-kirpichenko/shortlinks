package grpc

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetUserURLs -
func (s *Server) GetUserURLs(_, req *UserURLsRequest) (*UserURLsResponse, error) {

	var links []*URLsResponse
	userID := req.UserId

	if len(userID) == 0 {
		return nil, status.Error(codes.InvalidArgument, `user_id is empty!`)
	}

	userURLs, err := s.Storage.GetAllByUserID(userID)

	if err != nil {
		return nil, status.Error(codes.NotFound, "urls not found")
	}

	for _, url := range userURLs {
		resp := &URLsResponse{
			Short:    strings.TrimSpace(fmt.Sprintf(s.App.Cfg.ResultURL+"/%s", url.Short)),
			Original: strings.TrimSpace(url.Original),
		}

		links = append(links, resp)
	}
	return &UserURLsResponse{
		Urls: links,
	}, nil

}
