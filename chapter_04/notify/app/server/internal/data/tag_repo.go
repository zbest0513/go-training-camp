package data

import (
	"context"
	"fmt"
	guuid "github.com/google/uuid"
	"github.com/pkg/errors"
	"notify-server/internal/biz"
	"notify-server/internal/data/ent"
	tWhere "notify-server/internal/data/ent/tag"
	tlWhere "notify-server/internal/data/ent/templatetagrelation"
	uWhere "notify-server/internal/data/ent/user"
	utWhere "notify-server/internal/data/ent/usertagrelation"
	"time"
)

type tagRepo struct {
	data *ent.Client
}

func NewTagRepo(data *ent.Client) biz.TagRepo {
	return &tagRepo{
		data: data,
	}
}

func (tr *tagRepo) CreateTag(ctx context.Context, tag biz.Tag) (*biz.Tag, error) {
	now := time.Now()
	t, err := tr.data.Tag.Create().SetName(tag.Name).SetUUID(guuid.New().String()).
		SetDesc(tag.Desc).SetUpdatedAt(now).SetCreatedAt(now).Save(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "创建Tag失败")
	}

	result := &biz.Tag{
		Name:      t.Name,
		Desc:      t.Desc,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
		Status:    (int8)(t.Status),
		Uuid:      t.UUID,
		Id:        t.ID,
	}
	return result, nil
}

func (tr *tagRepo) QueryTagByName(ctx context.Context, name string) (*biz.Tag, error) {

	t, err := tr.data.Tag.Query().Where(tWhere.NameEQ(name)).Only(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "根据标签名称查询Tag失败")
	}
	result := &biz.Tag{
		Name:      t.Name,
		Desc:      t.Desc,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
		Status:    (int8)(t.Status),
		Uuid:      t.UUID,
		Id:        t.ID,
	}
	return result, nil
}

func (tr *tagRepo) SyncTag(ctx context.Context, tag biz.Tag) (int, error) {
	count, err := tr.data.Tag.Update().Where(tWhere.UUIDEQ(tag.Uuid)).SetDesc(tag.Desc).
		SetStatus((int)(tag.Status)).SetUpdatedAt(time.Now()).Save(ctx)
	if err != nil {
		return 0, errors.WithMessage(err, "repo更新Tag数据失败")
	}
	return count, nil
}

func (tr *tagRepo) UpdateStatus(ctx context.Context, uuid string, status int) (int, error) {
	count, err := tr.data.Tag.Update().Where(tWhere.UUIDEQ(uuid)).
		SetStatus(status).SetUpdatedAt(time.Now()).Save(ctx)
	if err != nil {
		return 0, errors.WithMessage(err, fmt.Sprintf("repo更新Tag状态%v失败", status))
	}
	return count, nil
}

func (tr *tagRepo) DeleteTag(ctx context.Context, uuid string) (int, error) {
	count, err := tr.data.Tag.Delete().Where(tWhere.UUIDEQ(uuid)).Exec(ctx)
	if err != nil {
		return 0, errors.WithMessage(err, "repo删除tag失败")
	}
	return count, nil
}

func (tr *tagRepo) DisbandUserRelations(ctx context.Context, uuid string) (int, error) {
	count, err := tr.data.UserTagRelation.Delete().Where(utWhere.TagUUIDEQ(uuid)).Exec(ctx)
	if err != nil {
		return count, errors.WithMessage(err, "删除用户关系失败")
	}
	return count, nil
}

func (tr *tagRepo) DisbandTemplateRelations(ctx context.Context, uuid string) (int, error) {
	count, err := tr.data.TemplateTagRelation.Delete().Where(tlWhere.TagUUIDEQ(uuid)).Exec(ctx)
	if err != nil {
		return count, errors.WithMessage(err, "删除模版关系失败")
	}
	return count, nil
}

func (tr *tagRepo) UpdateUserRelationsStatus(ctx context.Context, uuid string, status int) (int, error) {
	count, err := tr.data.UserTagRelation.Update().SetStatus(status).Where(utWhere.TagUUIDEQ(uuid)).Save(ctx)
	if err != nil {
		return count, errors.WithMessage(err, "修改用户关系状态失败")
	}
	return count, nil
}

func (tr *tagRepo) UpdateTemplateRelationsStatus(ctx context.Context, uuid string, status int) (int, error) {
	count, err := tr.data.TemplateTagRelation.Update().SetStatus(status).Where(tlWhere.TagUUIDEQ(uuid)).Save(ctx)
	if err != nil {
		return count, errors.WithMessage(err, "修改模版关系状态失败")
	}
	return count, nil
}

func (tr *tagRepo) QueryUsers(ctx context.Context, uuid string) ([]*biz.User, error) {
	all, err := tr.data.UserTagRelation.Query().Where(utWhere.TagUUIDEQ(uuid)).All(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("查询tag[%v]用户关系失败", uuid))
	}
	count := len(all)
	userUuids := make([]string, count)
	for i := 0; i < count; i++ {
		userUuids[i] = *all[i].UserUUID
	}
	users, err := tr.data.User.Query().Where(uWhere.UUIDIn(userUuids...)).All(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("查询用户列表失败:[%v]", userUuids))
	}
	count = len(users)
	result := make([]*biz.User, count)

	for i := 0; i < count; i++ {
		result[i] = &biz.User{
			users[i].ID,
			users[i].UUID,
			users[i].Name,
			users[i].Mobile,
			users[i].Email,
			(int8)(users[i].Status),
			users[i].CreatedAt,
			users[i].UpdatedAt,
		}
	}
	return result, nil
}
