package service

import (
	"context"

	"github.com/CharmingCharm/DouSheng/internal/video/db"
	"github.com/CharmingCharm/DouSheng/kitex_gen/video"
	"github.com/CharmingCharm/DouSheng/pkg/status"
)

type UpdateFavoriteCountService struct {
	ctx context.Context
}

// NewUpdateFavoriteCountService new UpdateFavoriteCountService
func NewUpdateFavoriteCountService(ctx context.Context) *UpdateFavoriteCountService {
	return &UpdateFavoriteCountService{ctx: ctx}
}

// Update the number of "likes"
func (s *UpdateFavoriteCountService) UpdateFavoriteCount(req *video.UpdateFavoriteCountRequest) error {
	if req.ActionType == 1 {
		err := db.VideoFavoriteCountAdd(req.VideoId)
		if err != nil {
			return err
		}
	} else if req.ActionType == 2 {
		err := db.VideoFavoriteCountSubtract(req.VideoId)
		if err != nil {
			return err
		}
	} else {
		return status.ParamErr
	}
	return nil
}
