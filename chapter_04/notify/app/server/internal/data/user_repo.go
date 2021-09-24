package data

import (
	"notify-server/internal/biz"
	"notify-server/internal/data/ent"
)

type userRepo struct {
	data *ent.Client
}

func NewUserRepo(data *ent.Client) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}
