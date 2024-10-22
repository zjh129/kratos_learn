package service

import (
	"context"

	"kratos_learn/api/user"
	"kratos_learn/internal/biz"
)

// UserService is a greeter service.
type UserService struct {
	user.UnimplementedUserServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *UserService {
	return &UserService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *UserService) UserFind(ctx context.Context, in *user.UserInfoRequest) (*user.UserInfo, error) {
	/*g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}*/
	return &user.UserInfo{
		Id:      in.Id,
		Name:    "",
		Email:   "",
		Phone:   "",
		Address: "",
	}, nil
}
