package model

type User struct {
	Id   int    `model:"id"`
	Name string `model:"name"`
	Age  int    `model:"age"`
	Card string `model:"card"`
}

func (receiver *User) New(user *User) *User {
	return &User{
		Id:   user.Id,
		Name: user.Name,
		Age:  user.Age,
		Card: user.Card,
	}
}
