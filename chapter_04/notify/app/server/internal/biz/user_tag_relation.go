package biz

import "time"

type UserTagRelation struct {
	Id        int
	UserUuid  string
	TagUuid   string
	Status    int8
	CreatedAt time.Time
	UpdatedAt time.Time
}
