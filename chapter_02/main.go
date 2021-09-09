package main

import (
	"chapter_02/model"
	"chapter_02/utils"
	"fmt"
	"log"
	"reflect"
)

func main() {
	user := model.User{Id: 12}

	elem := reflect.ValueOf(&user).Elem()
	value := elem.FieldByName("Age").Addr().Elem().Interface()

	fmt.Println(value)

	_, err := utils.QueryOne(&user)

	if err != nil {
		log.Println(fmt.Sprintf("%+v\n", err))
	}
	fmt.Println("==========")
}
