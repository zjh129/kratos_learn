package biz

import (
	"context"

	"kratos_learn/api/user"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(user.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// User is a User model.
type User struct {
	Id        int64
	Uqid      string
	Name      string
	Avatar    string
	Type      UserType
	IsEnable  bool
	Status    UserStatus
	CreatedAt string
	UpdatedAt string
}

type UserType int64

const (
	AdminUser  UserType = 1
	NormalUser UserType = 2
)

type UserStatus int64

const (
	NormalStatus  UserStatus = 1
	DisableStatus UserStatus = 2
)

// UserList is a User list.
type UserListCondition struct {
	Page     int64      // 页码
	PageSize int64      // 每页数量
	Name     string     // 名称
	Type     UserType   // 类型
	Status   UserStatus // 状态
}

// UserRepo is a User repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Delete(context.Context, int64) error
	FindByID(context.Context, int64) (*User, error)
	List(context.Context, *UserListCondition) ([]*User, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// SaveUser creates a User, and returns the new User.
func (uc *UserUsecase) SaveUser(ctx context.Context, u *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", u.Name)
	return uc.repo.Save(ctx, u)
}

// FindUser finds the User by the ID.
func (uc *UserUsecase) FindUser(ctx context.Context, id int64) (*User, error) {
	uc.log.WithContext(ctx).Infof("FindUser: %v", id)
	return uc.repo.FindByID(ctx, id)
}

// DeleteUser deletes the User by the ID.
func (uc *UserUsecase) DeleteUser(ctx context.Context, id int64) error {
	uc.log.WithContext(ctx).Infof("DeleteUser: %v", id)
	return uc.repo.Delete(ctx, id)
}

// ListUser lists the Users.
func (uc *UserUsecase) ListUser(ctx context.Context, c *UserListCondition) ([]*User, error) {
	uc.log.WithContext(ctx).Infof("ListUser: %v", c)
	return uc.repo.List(ctx, c)
}
