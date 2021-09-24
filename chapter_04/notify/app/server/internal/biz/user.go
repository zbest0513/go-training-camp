package biz

import "time"

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
