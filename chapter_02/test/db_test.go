package test

import (
	"chapter_02/model"
	"chapter_02/utils"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"testing"
)

func TestQueryOne(t *testing.T) {
	user := model.User{Name: "lisi"}
	where := new(utils.WhereGenerator).NewInstance().And(
		"name").Equals(user.Name).And("age").Equals(55)

	one, err := new(utils.DBUtils).QueryOne(&user, where, "card", "name", "age")
	if errors.Is(err, sql.ErrNoRows) {
		//panic(zerror.UserNotFound)
	} else if err != nil {
		log.Println(fmt.Sprintf("%+v", err))
	}
	log.Printf("result : %+v", one)
}

func TestQueryList(t *testing.T) {
	user := new(model.User)
	user.Age = 22
	where := new(utils.WhereGenerator).NewInstance().And("age").Equals(22)

	list, err, _ := new(utils.DBUtils).QueryList(user, where)
	if err != nil {
		log.Println(fmt.Sprintf("%+v\n", err))
	}
	for i, item := range list {
		log.Println(fmt.Sprintf("row : %+v , i:%v", item, i))
	}
}

func TestUpdateModel(t *testing.T) {
	user := new(model.User)
	user.Age = 23
	where := new(utils.WhereGenerator).NewInstance().And("age").Equals(12)
	count, err := new(utils.DBUtils).UpdateModels(user, where, []string{"age"})
	if err != nil {
		log.Println(fmt.Sprintf("处理异常....%+v", err))
	}
	log.Println(fmt.Sprintf("执行更新%v条成功", count))
}

func TestInsert(t *testing.T) {
	user := new(model.User)
	user.Age = 23
	user.Card = "33333333333"
	user.Name = "zzz"
	count, err := new(utils.DBUtils).InsertModels(user, user)
	if err != nil {
		log.Println(fmt.Sprintf("处理异常....%+v", err))
	}
	log.Println(fmt.Sprintf("执行插入%v条成功", count))
}

func TestDelete(t *testing.T) {
	user := new(model.User)
	user.Card = "33333333333"

	dbUtils := new(utils.DBUtils)
	where := new(utils.WhereGenerator).NewInstance().And("card").Equals(user.Card)
	count, err := dbUtils.DeleteModels(user, where)
	if err != nil {
		log.Println(fmt.Sprintf("处理异常....%+v", err))
	}
	log.Println(fmt.Sprintf("执行删除%v条成功", count))
}

func TestTrans(t *testing.T) {

	user := new(model.User)
	user.Age = 23
	user.Card = "444444444"
	user.Name = "kkkkk"

	db := new(utils.DBUtils)
	var insert = db.InsertModels
	executor := new(utils.TransTaskExecutor).NewInsertTaskExecutor(insert, user, user)

	db.TxExec(executor)

}
