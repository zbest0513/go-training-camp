package biz

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"notify-server/internal/pkg/enum"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Mobile    string
	Email     string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepo interface {
	Create(context.Context, User) (*User, error)
	QueryUserByMobile(context.Context, string) (*User, error)
	SyncUser(context.Context, User) (int, error)
	UpdateUserStatus(context.Context, string, int) (int, error)
	AddTags(ctx context.Context, userUuid string, tagUuids []string) (int, error)
	DeleteTags(ctx context.Context, userUuid string) (int, error)
	DisbandTags(ctx context.Context, userUuid string, tagUuids []string) (int, error)
	UpdateTagRelationsStatus(ctx context.Context, userUuid string, status int, tagUuids ...string) (int, error)
	QueryAll(ctx context.Context) ([]*User, error)
}

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(up UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: up,
	}
}

// CreateUser 创建用户
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

// UpdateUserStatus 启用/禁用用户
func (uc *UserUseCase) UpdateUserStatus(ctx context.Context, uuid string, status int) error {
	_, err := uc.repo.UpdateUserStatus(ctx, uuid, status)
	if err != nil {
		return errors.WithMessage(err, "UpdateUserStatus修改用户状态失败")
	}
	var rStatus int
	var logStr string
	if status == enum.USER_STATUS_AVAILABLE { //启用
		rStatus = enum.RELATION_USER_TAG_AVAILABLE
		logStr = "启用"
	} else { //禁用
		rStatus = enum.RELATION_USER_TAG_UNAVAILABLE
		logStr = "禁用"
	}
	count, err := uc.repo.UpdateTagRelationsStatus(ctx, uuid, rStatus)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("%v用户[%v]修改状态失败[%v]", logStr, uuid, count))
	}
	log.Println(fmt.Sprintf("%v用户[%v]修改状态成功[%v]", logStr, uuid, count))
	return nil
}

// AddTags 给用添加标签
func (uc *UserUseCase) AddTags(ctx context.Context, userUuid string, tagUuids []string) error {

	//删除关系
	count, err := uc.repo.DeleteTags(ctx, userUuid)
	if err != nil {
		return errors.WithMessage(err, "添加标签时,删除老的标签关系失败")
	}
	log.Println(fmt.Sprintf("添加标签时,删除老的标签关系成功[%v]", count))

	//重建关系
	count, err = uc.repo.AddTags(ctx, userUuid, tagUuids)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("给用户[%v]重建标签关系失败[%v]", userUuid, count))
	}
	log.Println(fmt.Sprintf("给用户[%v]重建标签关系成功[%v]", userUuid, count))
	return nil
}

// QueryByMobile 根据手机号查询用户
func (uc *UserUseCase) QueryByMobile(ctx context.Context, mobile string) (*User, error) {
	return uc.repo.QueryUserByMobile(ctx, mobile)
}

// QueryAll 查询用户列表
func (uc *UserUseCase) QueryAll(ctx context.Context) ([]*User, error) {
	all, err := uc.repo.QueryAll(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "biz查询用户列表失败")
	}
	return all, nil
}
