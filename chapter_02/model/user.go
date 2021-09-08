package model

type User struct {
	Id   int    `model:"id"`
	Name string `model:"name"`
	Age  int    `model:"age"`
}
