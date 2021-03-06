package data

import (
	"context"
	"fmt"
	guuid "github.com/google/uuid"
	xErrors "github.com/pkg/errors"
	"log"
	"notify-server/internal/biz"
	"notify-server/internal/data/ent"
	uWhere "notify-server/internal/data/ent/user"
	utWhere "notify-server/internal/data/ent/usertagrelation"
	"notify-server/internal/pkg/enum"
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
		return nil, xErrors.WithMessage(err, "创建用户失败")
	}
	return ur.user2DO(u), nil
}

func (ur *userRepo) QueryUserByMobile(ctx context.Context, mobile string) (*biz.User, error) {
	u, err := ur.data.User.Query().Where(uWhere.MobileEQ(mobile)).Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, nil
	} else if err != nil {
		return nil, xErrors.WithMessage(err, fmt.Sprintf("查询用户[%v]失败:", mobile))
	}
	return ur.user2DO(u), nil
}

func (ur *userRepo) SyncUser(ctx context.Context, user biz.User) (int, error) {
	count, err := ur.data.User.Update().Where(uWhere.UUIDEQ(user.Uuid)).SetEmail(user.Email).
		SetName(user.Name).SetUpdatedAt(time.Now()).SetStatus(0).Save(ctx)
	if err != nil {
		return 0, xErrors.WithMessage(err, fmt.Sprintf("修改用户失败:%v", user))
	}
	return count, nil
}

func (ur *userRepo) UpdateUserStatus(ctx context.Context, uuid string, status int) (int, error) {
	count, err := ur.data.User.Update().Where(uWhere.UUIDEQ(uuid)).SetStatus(status).Save(ctx)
	if err != nil {
		return 0, xErrors.WithMessage(err, fmt.Sprintf("修改用户状态失败:%v,%v", uuid, status))
	}
	return count, nil
}

func (ur *userRepo) DeleteTags(ctx context.Context, userUuid string) (int, error) {
	count, err := ur.data.UserTagRelation.Delete().Where(utWhere.UserUUIDEQ(userUuid)).Exec(ctx)
	if err != nil {
		return count, xErrors.WithMessage(err, fmt.Sprintf("删除用户[%v]的标签关系失败", userUuid))
	}
	return count, nil
}

func (ur *userRepo) AddTags(ctx context.Context, userUuid string, tagUuids []string) (int, error) {
	creates := make([]*ent.UserTagRelationCreate, len(tagUuids))
	now := time.Now()
	for i, tagUuid := range tagUuids {
		creates[i] = ur.data.UserTagRelation.Create().SetStatus(enum.RELATION_USER_TAG_AVAILABLE).
			SetCreatedAt(now).SetCreatedAt(now).SetTagUUID(tagUuid).SetUserUUID(userUuid)
	}
	relations, err := ur.data.UserTagRelation.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return 0, xErrors.WithMessage(err, fmt.Sprintf("批量建立用户[%v]的标签关系失败", userUuid))
	}
	return len(relations), nil
}

func (ur *userRepo) DisbandTags(ctx context.Context, userUuid string, tagUuids []string) (int, error) {

	count, err := ur.data.UserTagRelation.Delete().
		Where(utWhere.UserUUIDEQ(userUuid), utWhere.TagUUIDIn(tagUuids...)).Exec(ctx)
	if err != nil {
		return count, xErrors.WithMessage(err, fmt.Sprintf("解除用户[%v]的tags[%v]失败", userUuid, tagUuids))
	}
	log.Println(fmt.Sprintf("解除用户[%v]的tags[%v]成功[%v]", userUuid, tagUuids, count))
	return count, nil
}

func (ur *userRepo) UpdateTagRelationsStatus(ctx context.Context, userUuid string, status int, tagUuids ...string) (int, error) {
	update := ur.data.UserTagRelation.Update().SetStatus(status)
	if len(tagUuids) > 0 {
		update = update.Where(utWhere.UserUUIDEQ(userUuid), utWhere.TagUUIDIn(tagUuids...))
	} else {
		update = update.Where(utWhere.UserUUIDEQ(userUuid))
	}
	count, err := update.Save(ctx)
	if err != nil {
		return count, xErrors.WithMessage(err, fmt.Sprintf("修改用户[%v]的tags[%v]状态[%v]失败", userUuid, tagUuids, status))
	}
	log.Println(fmt.Sprintf("修改用户[%v]的tags[%v]的状态[%v]成功[%v]", userUuid, tagUuids, status, count))
	return count, nil
}

func (ur *userRepo) QueryAll(ctx context.Context) ([]*biz.User, error) {
	users, err := ur.data.User.Query().All(ctx)
	if err != nil {
		return nil, xErrors.WithMessage(err, "repo:查询用户列表失败")
	}
	result := make([]*biz.User, len(users))
	for i, user := range users {
		result[i] = ur.user2DO(user)
	}
	return result, nil
}

func (ur *userRepo) user2DO(user *ent.User) *biz.User {
	return &biz.User{
		Id:        user.ID,
		Uuid:      user.UUID,
		Name:      user.Name,
		Mobile:    user.Mobile,
		Email:     user.Email,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
