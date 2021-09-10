package main

import (
	"chapter_02/model"
	"chapter_02/utils"
	"fmt"
	"log"
)

func main() {
	userCase2()
}

func userCase1() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(fmt.Sprintf("%+v\n", r))
		}
	}()

	user := model.User{Name: "lisi"}
	where := new(utils.WhereGenerator).NewInstance().And("name").Equals(user.Name).And("age").Equals(12)

	one, err := utils.QueryOne(&user, where, "name")
	if err != nil {
		log.Println(fmt.Sprintf("%+v\n", err))
	}
	log.Fatalf("result : %+v", one)
}

func userCase2() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(fmt.Sprintf("%+v\n", r))
		}
	}()

	user := new(model.User)
	user.Age = 22
	where := new(utils.WhereGenerator).NewInstance().And("age").Equals(22)

	list, err, i := utils.QueryList(user, where)
	if err != nil {
		log.Println(fmt.Sprintf("%+v\n", err))
	}
	log.Fatalf("result : %+v , count:%v", list, i)
}
