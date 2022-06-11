package service

import (
	"context"

	"github.com/CharmingCharm/DouSheng/internal/video/db"
	"github.com/CharmingCharm/DouSheng/internal/video/rpc"
	"github.com/CharmingCharm/DouSheng/kitex_gen/action"
	"github.com/CharmingCharm/DouSheng/kitex_gen/base"
	"github.com/CharmingCharm/DouSheng/kitex_gen/user"
	"github.com/CharmingCharm/DouSheng/kitex_gen/video"
	"github.com/CharmingCharm/DouSheng/pkg/status"
)

type GetPublishedVideosService struct {
	ctx context.Context
}

// NewGetPublishedVideosService new GetPublishedVideosService
func NewGetPublishedVideosService(ctx context.Context) *GetPublishedVideosService {
	return &GetPublishedVideosService{ctx: ctx}
}

// Get the published videos
func (s *GetPublishedVideosService) GetPublishedVideos(req *video.GetPublishedVideosRequest) ([]*base.Video, error) {
	videoDBList, err := db.GetVideoListByAuthorId(req.UserId)
	if err != nil {
		return nil, err
	}
	var myId int64
	myId = req.MyId

	videoList := make([]*base.Video, len(videoDBList))

	for index, v := range videoDBList {
		userInfo, err := rpc.GetUserInfo(s.ctx, &user.GetUserInfoRequest{
			UserId: v.AuthId,
			MyId:   myId,
		})
		if err != nil {
			return nil, err
		}
		if userInfo.BaseResp.StatusCode != status.SuccessCode {
			return nil, status.NewStatus(userInfo.BaseResp.StatusCode, userInfo.BaseResp.StatusMessage)
		}

		flag := false
		if myId != -1 {
			favoriteInfo, err := rpc.CheckFavorite(s.ctx, &action.CheckFavoriteRequest{
				MyId:    myId,
				VideoId: v.Id,
			})
			if err != nil {
				return nil, err
			}
			if favoriteInfo.BaseResp.StatusCode != status.SuccessCode {
				return nil, status.NewStatus(favoriteInfo.BaseResp.StatusCode, favoriteInfo.BaseResp.StatusMessage)
			}
			flag = *favoriteInfo.IsFavorite
		}

		videoList[index] = &base.Video{
			Id:            v.Id,
			Author:        userInfo.User,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    flag,
			Title:         v.Title,
		}
	}

	return videoList, nil
}
