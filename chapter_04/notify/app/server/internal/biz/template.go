package biz

import (
	"context"
	"time"
)

type Template struct {
	Id        int
	Uuid      string
	Name      string
	Desc      string
	Content   string
	Status    int8
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TemplateRepo interface {
	CreateTemplate(context.Context, Template) (*Template, error)
	UpdateTemplate(context.Context, Template) (int, error)
	UpdateStatus(context.Context, string, int) (int, error)
	DeleteTemplate(context.Context, string) (int, error)
}
