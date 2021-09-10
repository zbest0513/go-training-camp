package utils

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

// 查询单条记录
// 参数说明:
//	target:条件载体，用户反射/填充等操作
//	where:where条件生成器,用于生成where条件
//	scan:需要检索的字段
// 返回说明:
//	interface:检索到的条目
//	error:异常信息

func QueryOne(target interface{}, where *WhereGenerator, scans ...string) (interface{}, error) {
	defer deferError("query one method error")
	//生成sql
	sql := queryOneSqlGenerate(target, where.Sql(), scans...)

	//获取要填充的字段
	values := getScanValues(target, scans...)
	//获取数据库连接
	db := GetConn()
	err := db.QueryRow(sql).Scan(values...)
	if err != nil {
		return nil, err
	}
	log.Println(fmt.Sprintf("query one result : %v", target))
	return target, nil
}

// 查询列表
// 参数说明:
//	target:条件载体，用户反射/填充等操作
//	where:where条件生成器,用于生成where条件
//	scan:需要检索的字段
// 返回说明:
//	[]interface:检索到的条目切片
//	error:异常信息

func QueryList(target interface{}, where *WhereGenerator, scans ...string) ([]interface{}, error, int) {
	defer deferError("query list method error")
	//生成sql
	sql := queryOneSqlGenerate(target, where.Sql(), scans...)
	//获取数据库连接
	db := GetConn()
	query, err := db.Query(sql)
	if err != nil {
		return nil, err, 0
	}
	defer query.Close()

	//默认20行，超过20行会动态扩容切片
	result := make([]interface{}, 20, 20)
	var count = 0
	for query.Next() {
		//每次填充需要不同的struct指针
		st := reflect.TypeOf(target).Elem()
		value := reflect.New(st)
		//获取要填充的字段
		values := getScanValues(&value, scans...)
		err = query.Scan(values...)
		//一条错误 不记录行数
		if err != nil {
			log.Println(fmt.Sprintf("扫描记录失败:%+v", err))
			continue
		}
		result[count] = st
		count++
	}
	err = query.Err()
	if err != nil {
		return nil, err, count
	}
	log.Println(fmt.Sprintf("query list result : %v", result[0:count]))
	return result[0:count], nil, count
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
	elem := reflect.TypeOf(target).Elem()
	num := len(scans)
	var flag bool
	if flag = num > 0; !flag {
		num = elem.NumField()
	}
	value := reflect.ValueOf(target).Elem()
	values := make([]interface{}, num, num)
	var tmp = 0
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i).Tag.Get("model")
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

// 单条查询sql生成
// 参数说明:
//	target : 要查询的实体指针，用于反射得到表名、表字段和struct属性的映射
//	where : 用户生成查询条件字符串
//	scans : 数据库字段名称的切片，用于指定查询/映射哪些字段，如果不传则查询全部字段
// 返回说明:
//	str : 返回生成的sql
//	err : 将errors 和panic 统一上抛处理（error wrap处理/panic 转 error）

func queryOneSqlGenerate(target interface{}, where string, scans ...string) string {
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
	sql := fmt.Sprintf(fmt.Sprintf("select %s from %s %s limit 1", join, split[len(split)-1], where), tags...)
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
