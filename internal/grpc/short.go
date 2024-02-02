package grpc

import (
	"errors"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/al-kirpichenko/shortlinks/internal/models"
	"github.com/al-kirpichenko/shortlinks/internal/services/keygen"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
)

// CreateShortURL -
func (s *Server) CreateShortURL(_, req *ShortURLRequest) (*ShortURLResponse, error) {

	url := req.Url

	if len(url) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Empty URL")
	}

	userID := req.UserId

	if len(userID) == 0 {
		userID = uuid.New().String()
	}

	link := &models.Link{
		Short:    keygen.GenerateKey(),
		Original: url,
		UserID:   userID,
	}

	if err := s.Storage.Insert(link); err != nil {
		if errors.Is(err, storage.ErrConflict) {
			link, err = s.Storage.GetShort(link.Original)
			if err != nil {
				zap.L().Error("Don't get short URL", zap.Error(err))
				return nil, status.Errorf(codes.Internal, "Internal error")
			}
		} else {
			zap.L().Error("Don't insert URL", zap.Error(err))
			return nil, status.Errorf(codes.Internal, "Internal error")
		}
	}
	return &ShortURLResponse{
		ShortUrl: link.Short,
	}, nil
}
