package test

import (
	"chapter_02/model"
	"chapter_02/utils"
	"fmt"
	"log"
	"testing"
)

func TestQueryOne(t *testing.T) {
	user := model.User{Name: "lisi"}
	where := new(utils.WhereGenerator).NewInstance().And(
		"name").Equals(user.Name).And("age").Equals(22)

	one, err := utils.QueryOne(&user, where, "name", "age")
	if err != nil {
		log.Println(fmt.Sprintf("%+v", err))
	}
	log.Printf("result : %+v", one)
}

func TestQueryList(t *testing.T) {
	user := new(model.User)
	user.Age = 22
	where := new(utils.WhereGenerator).NewInstance().And("age").Equals(22)

	list, err, _ := utils.QueryList(user, where)
	if err != nil {
		log.Println(fmt.Sprintf("%+v\n", err))
	}
	for i, item := range list {
		log.Println(fmt.Sprintf("row : %+v , i:%v", item, i))
	}
}
