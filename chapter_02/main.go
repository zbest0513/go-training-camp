package main

import (
	"chapter_02/model"
	"chapter_02/utils"
	"fmt"
	"log"
)

func main() {
	user := model.User{Name: "lisi"}

	generator := new(utils.WhereGenerator).NewInstance()
	where := generator.And("name").Equals(user.Name).Sql()
	fmt.Println(where)
	sql, err := utils.QueryOne(&user, where)
	fmt.Println(sql)

	if err != nil {
		log.Println(fmt.Sprintf("%+v\n", err))
	}
	fmt.Println("==========")
}
