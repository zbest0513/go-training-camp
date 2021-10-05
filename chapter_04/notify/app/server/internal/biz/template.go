package biz

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"notify-server/internal/pkg/enum"
	"sync"
	"time"
)

type Template struct {
	Id        int
	Uuid      string
	Name      string
	Desc      string
	Content   string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TemplateRepo interface {
	CreateTemplate(context.Context, Template) (*Template, error)
	UpdateTemplate(context.Context, Template) (int, error)
	UpdateStatus(context.Context, string, int) (int, error)
	DeleteTemplate(context.Context, string) (int, error)
	DisbandTags(ctx context.Context, uuid string) (int, error)
	UpdateTagRelationsStatus(ctx context.Context, uuid string, status int) (int, error)
	QueryTags(ctx context.Context, uuid string) ([]*Tag, error)
}

type TemplateUseCase struct {
	repo    TemplateRepo
	tagRepo TagRepo
}

func NewTemplateUseCase(repo TemplateRepo, tagRepo TagRepo) *TemplateUseCase {
	return &TemplateUseCase{
		repo:    repo,
		tagRepo: tagRepo,
	}
}

func (tu *TemplateUseCase) CreateTemplate(ctx context.Context, template Template) (*Template, error) {
	createTemplate, err := tu.repo.CreateTemplate(ctx, template)
	if err != nil {
		return nil, errors.WithMessage(err, "创建模版失败")
	}
	return createTemplate, nil
}

func (tu *TemplateUseCase) UpdateTemplate(ctx context.Context, template Template) error {
	_, err := tu.repo.UpdateTemplate(ctx, template)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("修改模版[%v]失败", template))
	}
	return nil
}

func (tu *TemplateUseCase) DeleteTemplate(ctx context.Context, uuid string) error {
	count, err := tu.repo.DeleteTemplate(ctx, uuid)

	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("删除模版[%v]失败", uuid))
	}
	if count > 0 {
		count, err = tu.repo.DisbandTags(ctx, uuid)
		if err != nil {
			return errors.WithMessage(err, fmt.Sprintf("删除标签时,解除tag关系失败"))
		}
	}
	log.Println(fmt.Sprintf("删除模版[%v]成功", uuid))
	return nil
}

func (tu *TemplateUseCase) UpdateTemplateStatus(ctx context.Context, uuid string, status int) error {
	count, err := tu.repo.UpdateStatus(ctx, uuid, status)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("更改模版[%v]状态[%v]", uuid, status))
	}
	if count > 0 {
		var tStatus int
		if status == enum.TEMPLATE_STATUS_AVAILABLE { //启用
			tStatus = enum.RELATION_TEMPLATE_TAG_AVAILABLE
		} else { //禁用
			tStatus = enum.RELATION_TEMPLATE_TAG_UNAVAILABLE
		}
		count, err = tu.repo.UpdateTagRelationsStatus(ctx, uuid, tStatus)
		if err != nil {
			return errors.WithMessage(err, fmt.Sprintf("解除模版[%v]的标签关系失败", uuid))
		}
	}
	return nil
}

func (tu *TemplateUseCase) QueryUsers(ctx context.Context, uuid string) ([]*User, error) {
	tags, err := tu.QueryTags(ctx, uuid)
	if err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("根据模版[%v]查询用户列表方法查询标签失败", uuid))
	}
	result := make([]*User, 0)
	m := make(map[string][]*User)
	//TODO 换掉wait group
	var wg sync.WaitGroup
	for _, tag := range tags {
		wg.Add(1)
		go func() {
			defer wg.Done()
			users, err := tu.tagRepo.QueryUsers(ctx, tag.Uuid)
			if err != nil {
			}
			m[tag.Name] = users
		}()
	}
	for _, mk := range m {
		result = append(result, mk...)
	}
	return result, nil
}

func (tu *TemplateUseCase) QueryTags(ctx context.Context, uuid string) ([]*Tag, error) {
	tags, err := tu.repo.QueryTags(ctx, uuid)
	if err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("根据模版[%v]查询标签失败", uuid))
	}
	return tags, nil
}
