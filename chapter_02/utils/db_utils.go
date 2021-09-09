package utils

import (
	"fmt"
	xerrors "github.com/pkg/errors"
	"reflect"
	"strings"
)

//查询单条记录
func QueryOne(target interface{}) (interface{}, error) {
	//生成sql
	sql, sql_err := queryOneSqlGenerate(target)

	if sql_err != nil {
		return nil, sql_err
	}
	//获取要填充的字段
	values := getScanValues(target)
	fmt.Printf("values:%v\n", values)
	//获取数据库连接
	db := GetConn()
	err := db.QueryRow(sql).Scan(values...)
	if err != nil {
		return nil, err
	}
	fmt.Printf("user:%v\n", target)
	return target, nil
}

func getScanValues(target interface{}) []interface{} {
	elem := reflect.TypeOf(target).Elem()
	num := elem.NumField()
	value := reflect.ValueOf(target).Elem()
	values := make([]interface{}, num, num)
	for i := 0; i < num; i++ {
		values[i] = value.FieldByName(elem.Field(i).Name).Addr().Interface()
	}
	return values
}

func queryOneSqlGenerate(target interface{}) (str string, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = xerrors.New(x)
			case error:
				err = xerrors.Wrap(x, "sql generate error ")
			default:
				err = xerrors.New("unbekannt panic")
			}
		}
	}()
	elem := reflect.TypeOf(target).Elem()
	name := strings.ToLower(reflect.TypeOf(target).String())
	split := strings.Split(name, ".")
	num := elem.NumField()
	tags := make([]interface{}, num, num)
	search := make([]string, num, num)
	for i := 0; i < num; i++ {
		tags[i] = elem.Field(i).Tag.Get("model")
		search[i] = "%v"
	}
	var i int = 0
	m := 1 / i
	fmt.Println(m)
	join := strings.Join(search, ",")
	sql := fmt.Sprintf(fmt.Sprintf("select %s from %s limit 1", join, split[len(split)-1]), tags...)
	return sql, nil
}
