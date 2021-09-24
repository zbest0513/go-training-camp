package data

import (
	"notify-server/internal/biz"
	"notify-server/internal/data/ent"
)

type templateRepo struct {
	data *ent.Client
}

func NewTemplateRepo(data *ent.Client) biz.TemplateRepo {
	return &templateRepo{
		data: data,
	}
}
