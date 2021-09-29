package biz

import (
	"context"
	"github.com/pkg/errors"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Mobile    string
	Email     string
	Status    int8
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepo interface {
	Create(context.Context, User) (*User, error)
	QueryUserByMobile(context.Context, string) (*User, error)
	SyncUser(context.Context, User) (int, error)
	UpdateUserStatus(context.Context, string, int) (int, error)
}

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(up UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: up,
	}
}
func (uc *UserUseCase) CreateUser(ctx context.Context, user User) (*User, error) {

	mobile, err := uc.repo.QueryUserByMobile(ctx, user.Mobile)
	if err != nil {
		return nil, errors.WithMessage(err, "创建用户方法异常")
	}
	if mobile == nil { //新增
		create, err := uc.repo.Create(ctx, user)
		if err != nil {
			return nil, errors.WithMessage(err, "创建用户失败")
		}
		return create, nil
	} else { //同步
		syncUser, err := uc.repo.SyncUser(ctx, user)
		if err != nil {
			return nil, errors.WithMessage(err, "同步用户失败")
		}
		if syncUser > 0 { //更新后重新查询
			mobile, err = uc.repo.QueryUserByMobile(ctx, user.Mobile)
			if err != nil {
				return nil, errors.WithMessage(err, "同步用户后查询方法异常")
			}
		}
		return mobile, nil
	}
}

func (uc *UserUseCase) UpdateUserStatus(ctx context.Context, user User) error {
	_, err := uc.repo.UpdateUserStatus(ctx, user.Uuid, int(user.Status))
	if err != nil {
		return errors.WithMessage(err, "UpdateUserStatus修改用户状态失败")
	}
	return nil
}
