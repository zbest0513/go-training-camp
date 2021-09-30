package data

import (
	"context"
	"fmt"
	guuid "github.com/google/uuid"
	"github.com/pkg/errors"
	"notify-server/internal/biz"
	"notify-server/internal/data/ent"
	tgWhere "notify-server/internal/data/ent/tag"
	tWhere "notify-server/internal/data/ent/template"
	ttWhere "notify-server/internal/data/ent/templatetagrelation"
	"time"
)

type templateRepo struct {
	data *ent.Client
}

func NewTemplateRepo(data *ent.Client) biz.TemplateRepo {
	return &templateRepo{
		data: data,
	}
}

func (tr *templateRepo) CreateTemplate(ctx context.Context, template biz.Template) (*biz.Template, error) {
	now := time.Now()
	t, err := tr.data.Template.Create().SetName(template.Name).
		SetUUID(guuid.New().String()).SetContent(template.Content).
		SetDesc(template.Desc).SetUpdatedAt(now).SetCreatedAt(now).Save(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "创建Template失败")
	}

	result := &biz.Template{
		Name:      t.Name,
		Desc:      t.Desc,
		Content:   t.Content,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
		Status:    (int8)(t.Status),
		Uuid:      t.UUID,
		Id:        t.ID,
	}
	return result, nil
}

func (tr *templateRepo) UpdateTemplate(ctx context.Context, template biz.Template) (int, error) {
	count, err := tr.data.Template.Update().Where(tWhere.UUIDEQ(template.Uuid)).
		SetDesc(template.Desc).SetContent(template.Content).SetUpdatedAt(time.Now()).Save(ctx)
	if err != nil {
		return 0, errors.WithMessage(err, "repo更新Tag数据失败")
	}
	return count, nil
}

func (tr *templateRepo) UpdateStatus(ctx context.Context, uuid string, status int) (int, error) {
	count, err := tr.data.Template.Update().Where(tWhere.UUIDEQ(uuid)).
		SetStatus(status).SetUpdatedAt(time.Now()).Save(ctx)
	if err != nil {
		return 0, errors.WithMessage(err, fmt.Sprintf("repo更新Tag状态%v失败", status))
	}
	return count, nil
}

func (tr *templateRepo) DeleteTemplate(ctx context.Context, uuid string) (int, error) {
	count, err := tr.data.Template.Delete().Where(tWhere.UUIDEQ(uuid)).Exec(ctx)
	if err != nil {
		return 0, errors.WithMessage(err, "repo删除tag失败")
	}
	return count, nil
}

func (tr *templateRepo) DisbandTags(ctx context.Context, uuid string) (int, error) {
	count, err := tr.data.TemplateTagRelation.Delete().Where(ttWhere.TemplateUUIDEQ(uuid)).Exec(ctx)
	if err != nil {
		return count, errors.WithMessage(err, fmt.Sprintf("删除template[%v]", uuid))
	}
	return count, nil
}

func (tr *templateRepo) UpdateTagRelationsStatus(ctx context.Context, uuid string, status int) (int, error) {
	count, err := tr.data.TemplateTagRelation.Update().SetStatus(status).Where(ttWhere.TemplateUUIDEQ(uuid)).Save(ctx)
	if err != nil {
		return count, errors.WithMessage(err, fmt.Sprintf("修改template[%v]的tag关系失败[%v]", uuid, status))
	}
	return count, nil
}

func (tr *templateRepo) QueryTags(ctx context.Context, uuid string) ([]*biz.Tag, error) {
	all, err := tr.data.TemplateTagRelation.Query().Where(ttWhere.TemplateUUIDEQ(uuid)).All(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("查询模版[%v]关联的tags失败", uuid))
	}
	ids := make([]int, len(all))
	for i, relation := range all {
		ids[i] = relation.ID
	}
	tags, err := tr.data.Tag.Query().Where(tgWhere.IDIn(ids...)).All(ctx)

	if err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("查询标签[%v]失败", ids))
	}
	result := make([]*biz.Tag, len(tags))
	for i, tag := range tags {
		result[i] = &biz.Tag{
			tag.ID,
			tag.UUID,
			tag.Name,
			tag.Desc,
			(int8)(tag.Status),
			tag.CreatedAt,
			tag.UpdatedAt,
		}
	}
	return result, nil
}
