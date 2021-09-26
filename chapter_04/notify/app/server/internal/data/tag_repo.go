package data

import (
	"notify-server/internal/biz"
	"notify-server/internal/data/ent"
)

type tagRepo struct {
	data *ent.Client
}

func NewTagRepo(data *ent.Client) biz.TagRepo {
	return &tagRepo{
		data: data,
	}
}
