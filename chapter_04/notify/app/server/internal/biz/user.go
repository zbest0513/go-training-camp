package biz

import (
	"context"
	"github.com/pkg/errors"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Mobile    string
	Email     string
	Status    int8
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepo interface {
	Create(context.Context, User) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(up UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: up,
	}
}
func (uc *UserUseCase) CreateUser(ctx context.Context, user User) (*User, error) {
	create, err := uc.repo.Create(ctx, user)
	if err != nil {
		return nil, errors.WithMessage(err, "创建用户service失败")
	}
	return create, nil
}
