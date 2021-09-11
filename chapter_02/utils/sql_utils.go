package utils

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

// 查询sql生成
// 参数说明:
//	target : 要查询的实体指针，用于反射得到表名、表字段和struct属性的映射
//	where : 用户生成查询条件字符串
//	scans : 数据库字段名称的切片，用于指定查询/映射哪些字段，如果不传则查询全部字段
// 返回说明:
//	str : 返回生成的sql
//	err : 将errors 和panic 统一上抛处理（error wrap处理/panic 转 error）

func querySqlGenerate(target interface{}, where string, scans ...string) string {
	defer deferError("query sql generate error")
	elem := reflect.TypeOf(target).Elem()
	name := strings.ToLower(reflect.TypeOf(target).String())
	split := strings.Split(name, ".")
	num := len(scans)
	var flag bool
	if flag = num > 0; !flag {
		num = elem.NumField()
	}
	tags := make([]interface{}, num, num)
	search := make([]string, num, num)
	var tmp = 0
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i).Tag.Get("model")
		if !flag || (flag && isExist(scans, field)) {
			tags[tmp] = field
			search[tmp] = "%v"
			tmp++
		}
	}
	join := strings.Join(search, ",")
	sql := fmt.Sprintf(fmt.Sprintf("select %s from %s %s ", join, split[len(split)-1], where), tags...)
	log.Println(fmt.Sprintf("生成sql:%v", sql))
	return sql
}

// 生成需要填充的属性的指针切片，用户查询结果回填struct
// 参数说明:
//	target : 要查询的实体指针，用于反射得到表名、表字段和struct属性的映射
//	scans : 数据库字段名称的切片，用于指定查询/映射哪些字段，如果不传则查询全部字段
// 返回说明:
//	values : 返回需要填充的属性的指针切片
//	err	:

func getScanValues(target interface{}, scans ...string) []interface{} {
	defer deferError("scan values reflect error")
	elem := reflect.TypeOf(target)
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}
	num := len(scans)
	var flag bool
	if flag = num > 0; !flag {
		num = elem.NumField()
	}
	value := reflect.ValueOf(target)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	values := make([]interface{}, num, num)
	var tmp = 0
	for i := 0; i < elem.NumField(); i++ {
		var field string
		field = elem.Field(i).Tag.Get("model")
		if !flag || (flag && isExist(scans, field)) {
			values[tmp] = value.FieldByName(elem.Field(i).Name).Addr().Interface()
			//将原有对象中的条件置为空，避免污染返回
			t := value.FieldByName(elem.Field(i).Name).Type()
			value.FieldByName(elem.Field(i).Name).Set(reflect.New(t).Elem())
			tmp++
		}
	}
	return values
}

// update sql生成
// 参数说明:
//	target : 要查询的实体指针，用于反射得到表名、表字段和struct属性的映射
//	where : 用户生成update条件字符串
//	sets : 数据库字段名称的切片，用于指定修改哪些字段，必传！
// 返回说明:
//	str : 返回生成的sql
//	err : 将errors 和panic 统一上抛处理（error wrap处理/panic 转 error）

func updateSqlGenerate(target interface{}, where string, sets []string) (string, []interface{}) {
	defer deferError(" update sql generate error")
	elem := reflect.TypeOf(target).Elem()
	name := strings.ToLower(reflect.TypeOf(target).String())
	split := strings.Split(name, ".")
	num := len(sets)
	var flag bool
	if flag = num > 0; !flag {
		num = elem.NumField()
	}
	tags := make([]interface{}, num, num)
	search := make([]string, num, num)
	values := make([]interface{}, num, num)
	value := reflect.ValueOf(target).Elem()
	var tmp = 0
	for i := 0; i < elem.NumField(); i++ {
		name := elem.Field(i).Name
		field := elem.Field(i).Tag.Get("model")
		if !flag || (flag && isExist(sets, field)) {
			tags[tmp] = field
			search[tmp] = fmt.Sprintf("%s = ? ", "%v")
			values[tmp] = value.FieldByName(name).Interface()
			tmp++
		}
	}
	join := strings.Join(search, ",")
	sql := fmt.Sprintf(fmt.Sprintf("update %s set %s %s ", split[len(split)-1], join, where), tags...)
	log.Println(fmt.Sprintf("生成sql:%v", sql))
	return sql, values
}

// insert sql生成
// 参数说明:
//	target : 要查询的实体指针，用于反射得到表名、表字段和struct属性的映射
// 返回说明:
//	str : 返回生成的sql
func insertSqlGenerate(models ...interface{}) string {
	defer deferError(" insert sql generate error")

	target := models[0]
	elem := reflect.TypeOf(target).Elem()
	name := strings.ToLower(reflect.TypeOf(target).String())
	split := strings.Split(name, ".")

	num := elem.NumField()
	tags := make([]string, num, num)
	values := make([]interface{}, num, num)
	value := reflect.ValueOf(target).Elem()
	var tmp = 0
	for i := 0; i < elem.NumField(); i++ {
		v := value.FieldByName(name)
		name := elem.Field(i).Name
		log.Println(fmt.Sprintf("============%v", v.IsNil()))
		if name == "Id" && v.CanInterface() == false {
			continue
		}
		field := elem.Field(i).Tag.Get("model")
		tags[tmp] = field
		values[tmp] = fmt.Sprintf("'%v'", v.Interface())
		tmp++
	}
	join := strings.Join(tags, ",")
	sql := fmt.Sprintf(fmt.Sprintf("insert into %s (%s) values (%s) ", split[len(split)-1], join, values))
	log.Println(fmt.Sprintf("生成sql:%v", sql))
	return sql
}

func isExist(fields []string, target string) bool {
	for _, field := range fields {
		if field == target {
			return true
		}
	}
	return false
}
