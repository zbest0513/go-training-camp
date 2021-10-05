package biz

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"notify-server/internal/pkg/enum"
	"time"
)

type Tag struct {
	Id        int
	Uuid      string
	Name      string
	Desc      string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
type TemplateTagRelation struct {
	Id           int
	TemplateUuid string
	TagUuid      string
	Status       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserTagRelation struct {
	Id        int
	UserUuid  string
	TagUuid   string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TagRepo interface {
	CreateTag(ctx context.Context, tag Tag) (*Tag, error)
	QueryTagByName(ctx context.Context, name string) (*Tag, error)
	SyncTag(ctx context.Context, tag Tag) (int, error)
	UpdateStatus(ctx context.Context, uuid string, status int) (int, error)
	DeleteTag(ctx context.Context, uuid string) (int, error)
	DisbandUserRelations(ctx context.Context, uuid string) (int, error)
	DisbandTemplateRelations(ctx context.Context, uuid string) (int, error)
	UpdateUserRelationsStatus(ctx context.Context, uuid string, status int) (int, error)
	UpdateTemplateRelationsStatus(ctx context.Context, uuid string, status int) (int, error)
	QueryUsers(ctx context.Context, uuid string) ([]*User, error)
}

type TagUseCase struct {
	repo TagRepo
}

func NewTagUseCase(repo TagRepo) *TagUseCase {
	return &TagUseCase{
		repo: repo,
	}
}

func (nt *TagUseCase) CreateTag(ctx context.Context, tag Tag) (*Tag, error) {
	result, err := nt.repo.CreateTag(ctx, tag)
	if err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("创建Tag[%v]失败", tag))
	}
	return result, nil
}

func (nt *TagUseCase) UpdateDesc(ctx context.Context, tag Tag) error {
	_, err := nt.repo.SyncTag(ctx, tag)
	if err != nil {
		return err
	}
	return nil
}

func (nt *TagUseCase) UpdateTagStatus(ctx context.Context, uuid string, status int) error {
	count, err := nt.repo.UpdateStatus(ctx, uuid, status)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("修改tag[%v]状态[%v]失败", uuid, status))
	}
	if count < 1 {
		log.Println(fmt.Sprintf("tag[%v]修改状态[%v]未生效", uuid, status))
		return nil
	}
	var rUserStatus int
	var rTemplateStatus int
	if status == enum.TAG_STATUS_AVAILABLE { //启用
		rUserStatus = enum.RELATION_USER_TAG_AVAILABLE
		rTemplateStatus = enum.RELATION_TEMPLATE_TAG_AVAILABLE
	} else { //禁用
		rUserStatus = enum.RELATION_USER_TAG_UNAVAILABLE
		rTemplateStatus = enum.RELATION_TEMPLATE_TAG_UNAVAILABLE
	}
	count, err = nt.repo.UpdateUserRelationsStatus(ctx, uuid, rUserStatus)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("修改tag[%v]用户关系状态[%v]失败", uuid, rUserStatus))
	}
	count, err = nt.repo.UpdateTemplateRelationsStatus(ctx, uuid, rTemplateStatus)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("修改tag[%v]模版关系状态[%v]失败", uuid, rTemplateStatus))
	}
	log.Println(fmt.Sprintf("tag[%v]修改状态[%v]成功", uuid, status))
	return nil
}

func (nt *TagUseCase) DeleteTag(ctx context.Context, uuid string) error {
	count, err := nt.repo.DeleteTag(ctx, uuid)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("删除tag[%v]失败", uuid))
	}
	if count < 1 {
		log.Println(fmt.Sprintf("tag[%v]删除未生效", uuid))
		return nil
	}
	count, err = nt.repo.DisbandUserRelations(ctx, uuid)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("解除tag[%v]用户关系失败", uuid))
	}
	count, err = nt.repo.DisbandTemplateRelations(ctx, uuid)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("解除tag[%v]模版关系失败", uuid))
	}
	log.Println(fmt.Sprintf("tag[%v]删除成功", uuid))
	return nil
}
