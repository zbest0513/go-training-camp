package biz

import "time"

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
}
