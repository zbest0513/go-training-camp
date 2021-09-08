package main

import (
	"chapter_02/model"
	"fmt"
	"reflect"
)

func main() {

	user := model.User{Id: 12}

	elem := reflect.ValueOf(&user).Elem()
	value := elem.FieldByName("Age").Addr().Elem().Interface()

	switch v := value.(type) {
	case int:
		println("int 类型", v == 0)
	case string:
		println("string 类型", v == "")
	default:
		println("``````")
	}

	fmt.Println(value)

	//one, err := utils.QueryOne(&user)
	//if err != nil {
	//	fmt.Printf("错误:%v\n",err)
	//}
	//fmt.Printf("查询的结果:%v",one)
}
