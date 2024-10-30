package service

import (
	"context"

	"kratos_learn/api/user"
	"kratos_learn/internal/biz"
)

// UserService is a greeter service.
type UserService struct {
	user.UnimplementedUserServer

	uc *biz.UserUsecase
}

// NewUserService new a greeter service.
func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

// UserSave implements user.UserServer.UserSave.
func (s *UserService) UserSave(ctx context.Context, in *user.UserSaveRequest) (*user.UserSaveReply, error) {
	u, err := s.uc.SaveUser(ctx, &biz.User{
		Id:       in.Id,
		Name:     in.Name,
		Avatar:   in.Avatar,
		Type:     biz.UserType(in.Type),
		IsEnable: in.IsEnable,
		Status:   biz.UserStatus(in.Status),
	})
	if err != nil {
		return nil, err
	}

	return &user.UserSaveReply{
		Id: u.Id,
	}, nil
}

// UserFind implements user.UserServer.
func (s *UserService) UserFind(ctx context.Context, in *user.UserInfoRequest) (*user.UserInfo, error) {
	u, err := s.uc.FindUser(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &user.UserInfo{
		Id:        u.Id,
		Uqid:      u.Uqid,
		Name:      u.Name,
		Avatar:    u.Avatar,
		Type:      user.UserType(u.Type),
		IsEnable:  u.IsEnable,
		Status:    user.UserStatus(u.Status),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

// UserList implements user.UserServer.UserList
func (s *UserService) UserList(ctx context.Context, in *user.UserListRequest) (*user.UserListReply, error) {
	ul, err := s.uc.ListUser(ctx, &biz.UserListCondition{
		Page:     in.Page,
		PageSize: in.PageSize,
		Name:     in.Name,
		Type:     biz.UserType(in.Type),
	})
	if err != nil {
		return nil, err
	}
	list := make([]*user.UserInfo, 0)
	for _, u := range ul {
		list = append(list, &user.UserInfo{
			Id:        u.Id,
			Uqid:      u.Uqid,
			Name:      u.Name,
			Avatar:    u.Avatar,
			Type:      user.UserType(u.Type),
			IsEnable:  u.IsEnable,
			Status:    user.UserStatus(u.Status),
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}
	return &user.UserListReply{List: list}, nil
}

// UserDelete implements user.UserServer.UserDelete
func (s *UserService) UserDelete(ctx context.Context, in *user.UserDeleteRequest) (*user.UserDeleteReply, error) {
	err := s.uc.DeleteUser(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &user.UserDeleteReply{
		Id: in.Id,
	}, nil
}
