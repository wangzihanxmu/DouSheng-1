// Code generated by Kitex v0.3.1. DO NOT EDIT.

package actionservice

import (
	"context"
	"github.com/CharmingCharm/DouSheng/idl/kitex_gen/action"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return actionServiceServiceInfo
}

var actionServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ActionService"
	handlerType := (*action.ActionService)(nil)
	methods := map[string]kitex.MethodInfo{
		"updateFavorite":      kitex.NewMethodInfo(updateFavoriteHandler, newActionServiceUpdateFavoriteArgs, newActionServiceUpdateFavoriteResult, false),
		"getFavoriteVideos":   kitex.NewMethodInfo(getFavoriteVideosHandler, newActionServiceGetFavoriteVideosArgs, newActionServiceGetFavoriteVideosResult, false),
		"updateComment":       kitex.NewMethodInfo(updateCommentHandler, newActionServiceUpdateCommentArgs, newActionServiceUpdateCommentResult, false),
		"getCommentLists":     kitex.NewMethodInfo(getCommentListsHandler, newActionServiceGetCommentListsArgs, newActionServiceGetCommentListsResult, false),
		"updateRelationship":  kitex.NewMethodInfo(updateRelationshipHandler, newActionServiceUpdateRelationshipArgs, newActionServiceUpdateRelationshipResult, false),
		"getUserFollowList":   kitex.NewMethodInfo(getUserFollowListHandler, newActionServiceGetUserFollowListArgs, newActionServiceGetUserFollowListResult, false),
		"getUserFollowerList": kitex.NewMethodInfo(getUserFollowerListHandler, newActionServiceGetUserFollowerListArgs, newActionServiceGetUserFollowerListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "action",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.3.1",
		Extra:           extra,
	}
	return svcInfo
}

func updateFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*action.ActionServiceUpdateFavoriteArgs)
	realResult := result.(*action.ActionServiceUpdateFavoriteResult)
	success, err := handler.(action.ActionService).UpdateFavorite(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newActionServiceUpdateFavoriteArgs() interface{} {
	return action.NewActionServiceUpdateFavoriteArgs()
}

func newActionServiceUpdateFavoriteResult() interface{} {
	return action.NewActionServiceUpdateFavoriteResult()
}

func getFavoriteVideosHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*action.ActionServiceGetFavoriteVideosArgs)
	realResult := result.(*action.ActionServiceGetFavoriteVideosResult)
	success, err := handler.(action.ActionService).GetFavoriteVideos(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newActionServiceGetFavoriteVideosArgs() interface{} {
	return action.NewActionServiceGetFavoriteVideosArgs()
}

func newActionServiceGetFavoriteVideosResult() interface{} {
	return action.NewActionServiceGetFavoriteVideosResult()
}

func updateCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*action.ActionServiceUpdateCommentArgs)
	realResult := result.(*action.ActionServiceUpdateCommentResult)
	success, err := handler.(action.ActionService).UpdateComment(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newActionServiceUpdateCommentArgs() interface{} {
	return action.NewActionServiceUpdateCommentArgs()
}

func newActionServiceUpdateCommentResult() interface{} {
	return action.NewActionServiceUpdateCommentResult()
}

func getCommentListsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*action.ActionServiceGetCommentListsArgs)
	realResult := result.(*action.ActionServiceGetCommentListsResult)
	success, err := handler.(action.ActionService).GetCommentLists(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newActionServiceGetCommentListsArgs() interface{} {
	return action.NewActionServiceGetCommentListsArgs()
}

func newActionServiceGetCommentListsResult() interface{} {
	return action.NewActionServiceGetCommentListsResult()
}

func updateRelationshipHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*action.ActionServiceUpdateRelationshipArgs)
	realResult := result.(*action.ActionServiceUpdateRelationshipResult)
	success, err := handler.(action.ActionService).UpdateRelationship(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newActionServiceUpdateRelationshipArgs() interface{} {
	return action.NewActionServiceUpdateRelationshipArgs()
}

func newActionServiceUpdateRelationshipResult() interface{} {
	return action.NewActionServiceUpdateRelationshipResult()
}

func getUserFollowListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*action.ActionServiceGetUserFollowListArgs)
	realResult := result.(*action.ActionServiceGetUserFollowListResult)
	success, err := handler.(action.ActionService).GetUserFollowList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newActionServiceGetUserFollowListArgs() interface{} {
	return action.NewActionServiceGetUserFollowListArgs()
}

func newActionServiceGetUserFollowListResult() interface{} {
	return action.NewActionServiceGetUserFollowListResult()
}

func getUserFollowerListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*action.ActionServiceGetUserFollowerListArgs)
	realResult := result.(*action.ActionServiceGetUserFollowerListResult)
	success, err := handler.(action.ActionService).GetUserFollowerList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newActionServiceGetUserFollowerListArgs() interface{} {
	return action.NewActionServiceGetUserFollowerListArgs()
}

func newActionServiceGetUserFollowerListResult() interface{} {
	return action.NewActionServiceGetUserFollowerListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) UpdateFavorite(ctx context.Context, req *action.UpdateFavoriteRequest) (r *action.UpdateFavoriteResponse, err error) {
	var _args action.ActionServiceUpdateFavoriteArgs
	_args.Req = req
	var _result action.ActionServiceUpdateFavoriteResult
	if err = p.c.Call(ctx, "updateFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFavoriteVideos(ctx context.Context, req *action.GetFavoriteVideosRequest) (r *action.GetFavoriteVideosResponse, err error) {
	var _args action.ActionServiceGetFavoriteVideosArgs
	_args.Req = req
	var _result action.ActionServiceGetFavoriteVideosResult
	if err = p.c.Call(ctx, "getFavoriteVideos", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateComment(ctx context.Context, req *action.UpdateCommentRequest) (r *action.UpdateCommentResponse, err error) {
	var _args action.ActionServiceUpdateCommentArgs
	_args.Req = req
	var _result action.ActionServiceUpdateCommentResult
	if err = p.c.Call(ctx, "updateComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCommentLists(ctx context.Context, req *action.GetCommentListsRequest) (r *action.GetCommentListsResponse, err error) {
	var _args action.ActionServiceGetCommentListsArgs
	_args.Req = req
	var _result action.ActionServiceGetCommentListsResult
	if err = p.c.Call(ctx, "getCommentLists", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateRelationship(ctx context.Context, req *action.UpdateRelationshipRequest) (r *action.UpdateRelationshipResponse, err error) {
	var _args action.ActionServiceUpdateRelationshipArgs
	_args.Req = req
	var _result action.ActionServiceUpdateRelationshipResult
	if err = p.c.Call(ctx, "updateRelationship", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserFollowList(ctx context.Context, req *action.GetUserFollowListRequest) (r *action.GetUserFollowListResponse, err error) {
	var _args action.ActionServiceGetUserFollowListArgs
	_args.Req = req
	var _result action.ActionServiceGetUserFollowListResult
	if err = p.c.Call(ctx, "getUserFollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserFollowerList(ctx context.Context, req *action.GetUserFollowerListRequest) (r *action.GetUserFollowerListResponse, err error) {
	var _args action.ActionServiceGetUserFollowerListArgs
	_args.Req = req
	var _result action.ActionServiceGetUserFollowerListResult
	if err = p.c.Call(ctx, "getUserFollowerList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
