package biz

import (
	"context"
	"time"
)

type Tag struct {
	Id        int
	Uuid      string
	Name      string
	Desc      string
	Status    int8
	CreatedAt time.Time
	UpdatedAt time.Time
}
type TemplateTagRelation struct {
	Id           int
	TemplateUuid string
	TagUuid      string
	Status       int8
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserTagRelation struct {
	Id        int
	UserUuid  string
	TagUuid   string
	Status    int8
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TagRepo interface {
	CreateTag(ctx context.Context, tag Tag) (*Tag, error)
	QueryTagByName(ctx context.Context, name string) (*Tag, error)
	SyncTag(ctx context.Context, tag Tag) (int, error)
	UpdateStatus(ctx context.Context, uuid string, status int) (int, error)
	DeleteTag(ctx context.Context, uuid string) (int, error)
}
