package data

import (
	"context"
	"fmt"
	guuid "github.com/google/uuid"
	"github.com/pkg/errors"
	"notify-server/internal/biz"
	"notify-server/internal/data/ent"
	uWhere "notify-server/internal/data/ent/user"
	"time"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *ent.Client
}

func NewUserRepo(data *ent.Client) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

func (ur *userRepo) Create(ctx context.Context, user biz.User) (*biz.User, error) {
	uuid := guuid.New().String()
	u, err := ur.data.User.Create().SetUUID(uuid).SetEmail(user.Email).
		SetMobile(user.Mobile).SetName(user.Name).Save(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "创建用户失败")
	}
	result := new(biz.User)
	result.Id = u.ID
	result.Name = u.Name
	result.Email = u.Email
	result.Mobile = u.Mobile
	result.CreatedAt = u.CreatedAt
	result.UpdatedAt = u.UpdatedAt
	result.Uuid = u.UUID
	result.Status = int8(u.Status)
	return result, nil
}

func (ur *userRepo) QueryUserByMobile(ctx context.Context, mobile string) (*biz.User, error) {
	u, err := ur.data.User.Query().Where(uWhere.MobileEQ(mobile)).Only(ctx)

	if err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("查询用户[%v]失败:", mobile))
	}
	result := new(biz.User)
	result.Id = u.ID
	result.Name = u.Name
	result.Email = u.Email
	result.Mobile = u.Mobile
	result.CreatedAt = u.CreatedAt
	result.UpdatedAt = u.UpdatedAt
	result.Status = int8(u.Status)
	return result, nil
}

func (ur *userRepo) SyncUser(ctx context.Context, user biz.User) (int, error) {
	count, err := ur.data.User.Update().Where(uWhere.UUIDEQ(user.Uuid)).SetEmail(user.Email).
		SetName(user.Name).SetUpdatedAt(time.Now()).SetStatus(0).Save(ctx)
	if err != nil {
		return 0, errors.WithMessage(err, fmt.Sprintf("修改用户失败:%v", user))
	}
	return count, nil
}

func (ur *userRepo) UpdateUserStatus(ctx context.Context, uuid string, status int) (int, error) {
	count, err := ur.data.User.Update().Where(uWhere.UUIDEQ(uuid)).SetStatus(status).Save(ctx)
	if err != nil {
		return 0, errors.WithMessage(err, fmt.Sprintf("修改用户状态失败:%v,%v", uuid, status))
	}
	return count, nil
}
