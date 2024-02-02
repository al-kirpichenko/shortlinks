package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/al-kirpichenko/shortlinks/internal/services/delurls"
)

// DeleteURLs -
func (s *Server) DeleteURLs(_, req *DeleteURLsRequest) (*DeleteURLsResponse, error) {

	userID := req.UserId

	if len(userID) == 0 {
		return nil, status.Error(codes.InvalidArgument, `user_id is empty!`)
	}

	urls := req.GetShorts()

	s.App.Worker.Channel <- &delurls.Task{
		UserID: userID,
		URLs:   urls,
	}
	return &DeleteURLsResponse{
		Result: "OK",
	}, nil
}
