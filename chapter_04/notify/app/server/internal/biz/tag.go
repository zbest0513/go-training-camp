package biz

import "time"

type Tag struct {
	Id        int
	Uuid      string
	Name      string
	Desc      string
	Status    int8
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TagRepo interface {
}
