package service

import (
	"context"

	// "github.com/CharmingCharm/DouSheng/internal/action/rpc"

	"github.com/CharmingCharm/DouSheng/internal/action/db"
	"github.com/CharmingCharm/DouSheng/kitex_gen/action"
)

type CheckRelationService struct {
	ctx context.Context
}

// NewCheckRelationService new CheckRelationService
func NewCheckRelationService(ctx context.Context) *CheckRelationService {
	return &CheckRelationService{ctx: ctx}
}

// CreateUser create user info.
func (s *CheckRelationService) CheckRelation(req *action.CheckRelationRequest) (bool, error) {

	// type CheckRelationRequest struct {
	// 	MyId    int64 `thrift:"my_id,1,required" json:"my_id"`
	// 	UserId  int64 `thrift:"u_id,2,required" json:"user_id"`
	// }
	flag, err := db.FindRelationRecord(s.ctx, req.MyId, req.UserId)
	if err != nil {
		return false, err
	}
	return flag, nil
}
