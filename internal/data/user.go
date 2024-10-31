package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	cache "github.com/mgtv-tech/jetcache-go"
	"kratos_learn/internal/biz"
	"kratos_learn/internal/data/ent"
	"kratos_learn/internal/data/ent/user"
	"time"
)

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
	}
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (u *userRepo) getCacheKey(id int64) string {
	return fmt.Sprintf("kratos_learn:user:%d", id)
}

func (u *userRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	var err error
	var get *ent.User
	if user.Id > 0 {
		get, err = u.data.db.User.Get(ctx, uint32(user.Id))
		if err != nil {
			return nil, err
		}
		ur := get.Update().
			SetName(user.Name).
			SetAvatar(user.Avatar).
			SetType(uint8(user.Type)).
			SetStatus(uint8(user.Status))
		if user.IsEnable {
			ur.SetIsEnable(1)
		} else {
			ur.SetIsEnable(0)
		}
		get, err = ur.SetUpdatedAt(time.Now()).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		uuidV7, _ := uuid.NewV7()
		cr := u.data.db.User.Create().
			SetUqid(uuidV7.String()).
			SetName(user.Name).
			SetAvatar(user.Avatar).
			SetType(uint8(user.Type)).
			SetStatus(uint8(user.Status))
		if user.IsEnable {
			cr.SetIsEnable(1)
		} else {
			cr.SetIsEnable(0)
		}
		get, err = cr.SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}
	// cache
	cacheKey := u.getCacheKey(int64(get.ID))
	err = u.data.cache.Set(ctx, cacheKey, cache.Value(get), cache.TTL(time.Hour))
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:        int64(get.ID),
		Uqid:      get.Uqid,
		Name:      get.Name,
		Avatar:    get.Avatar,
		Type:      biz.UserType(get.Type),
		IsEnable:  get.IsEnable == 1,
		Status:    biz.UserStatus(get.Status),
		CreatedAt: get.CreatedAt.Format(time.DateTime),
		UpdatedAt: get.UpdatedAt.Format(time.DateTime),
	}, nil
}

func (u *userRepo) Delete(ctx context.Context, i int64) error {
	err := u.data.db.User.DeleteOneID(uint32(i)).Exec(ctx)
	if err != nil {
		return err
	}
	cacheKey := u.getCacheKey(i)
	err = u.data.cache.Delete(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) FindByID(ctx context.Context, i int64) (*biz.User, error) {
	cacheKey := u.getCacheKey(i)
	getu := &ent.User{}
	if err := u.data.cache.Once(ctx, cacheKey, cache.Value(getu), cache.TTL(time.Hour), cache.Refresh(true),
		cache.Do(func(ctx context.Context) (any, error) {
			return u.data.db.User.Get(ctx, uint32(i))
		})); err != nil {
		return nil, err
	}
	return &biz.User{
		Id:        int64(getu.ID),
		Uqid:      getu.Uqid,
		Name:      getu.Name,
		Avatar:    getu.Avatar,
		Type:      biz.UserType(getu.Type),
		IsEnable:  getu.IsEnable == 1,
		Status:    biz.UserStatus(getu.Status),
		CreatedAt: getu.CreatedAt.Format(time.DateTime),
		UpdatedAt: getu.UpdatedAt.Format(time.DateTime),
	}, nil
}

// List .
func (u *userRepo) List(ctx context.Context, condition *biz.UserListCondition) ([]*biz.User, error) {
	query := u.data.db.User.Query()
	if condition != nil {
		if condition.Name != "" {
			query.Where(user.Name(condition.Name))
		}
		if condition.Type != 0 {
			query.Where(user.Type(uint8(condition.Type)))
		}
		if condition.Status != 0 {
			query.Where(user.Status(uint8(condition.Status)))
		}
		if condition.Page > 0 && condition.PageSize > 0 {
			query.Offset(int((condition.Page - 1) * condition.PageSize)).Limit(int(condition.PageSize))
		}
	}
	users, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	var result []*biz.User
	for _, user := range users {
		result = append(result, &biz.User{
			Id:        int64(user.ID),
			Uqid:      user.Uqid,
			Name:      user.Name,
			Avatar:    user.Avatar,
			Type:      biz.UserType(user.Type),
			IsEnable:  user.IsEnable == 1,
			Status:    biz.UserStatus(user.Status),
			CreatedAt: user.CreatedAt.Format(time.DateTime),
			UpdatedAt: user.UpdatedAt.Format(time.DateTime),
		})
	}
	return result, nil
}
