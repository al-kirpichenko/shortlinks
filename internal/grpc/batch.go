package grpc

import (
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/al-kirpichenko/shortlinks/internal/models"
	"github.com/al-kirpichenko/shortlinks/internal/services/keygen"
)

// ShortURLBatch -
func (s *Server) ShortURLBatch(_, req *ShortURLBatchRequest) (*ShortURLBatchResponse, error) {

	var (
		response []*BatchResponse
		links    []*models.Link
		userID   string
	)

	userID = req.UserId

	if userID == "" {
		userID = uuid.New().String()
	}

	for _, url := range req.GetUrls() {

		key := keygen.GenerateKey()
		resp := &BatchResponse{
			CorrelationId: url.CorrelationId,
			ResultUrl:     fmt.Sprintf(s.App.Cfg.ResultURL+"/%s", key),
		}
		link := &models.Link{
			Short:    key,
			Original: url.OriginalUrl,
			UserID:   userID,
		}

		response = append(response, resp)
		links = append(links, link)

	}

	if err := s.Storage.InsertLinks(links); err != nil {
		return nil, status.Error(codes.InvalidArgument, `Don't insert URLs`)
	}

	return &ShortURLBatchResponse{
		Urls: response,
	}, nil
}
