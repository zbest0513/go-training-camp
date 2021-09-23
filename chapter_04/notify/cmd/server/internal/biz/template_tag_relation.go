package biz

import "time"

type TemplateTagRelation struct {
	Id           int
	TemplateUuid string
	TagUuid      string
	Status       int8
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
