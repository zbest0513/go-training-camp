package utils

import (
	"fmt"
	"log"
	"reflect"
)

//QueryOne
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
	sql := querySqlGenerate(target, where.Sql(), scans...)

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

//QueryList
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
	sql := querySqlGenerate(target, where.Sql(), scans...)
	//获取数据库连接
	db := GetConn()
	query, err := db.Query(sql)
	if err != nil {
		return nil, err, 0
	}
	defer func() {
		e := query.Close()
		log.Println(fmt.Sprintf("db query close error %+v", e))
	}()

	//默认20行，超过20行会动态扩容切片
	result := make([]interface{}, 20, 20)
	var count = 0
	for query.Next() {
		//获取要填充的字段
		values := getScanValues(target, scans...)
		err = query.Scan(values...)
		//一条错误 不记录行数
		if err != nil {
			log.Println(fmt.Sprintf("扫描记录失败:%+v", err))
			continue
		}
		//每次填充需要不同的struct指针
		st := reflect.ValueOf(target)
		param := make([]reflect.Value, 1, 1)
		param[0] = reflect.ValueOf(target)
		calls := st.MethodByName("New").Call(param)
		value := calls[0]
		result[count] = value
		count++
	}
	err = query.Err()
	if err != nil {
		return nil, err, count
	}
	log.Println(fmt.Sprintf("query list result : %s,count : %v", result[0:count], count))
	return result[0:count], nil, count
}

// UpdateModels
// 查询列表
// 参数说明:
//	target:条件载体，用户反射/填充等操作
//	where:where条件生成器,用于生成where条件
//	sets:需要修改的字段,必传
// 返回说明:
//	int64:更新生效的条目数量
//	error:异常信息
func UpdateModels(target interface{}, where *WhereGenerator, sets []string) (int64, error) {
	defer deferError("update model method error")
	generate, values := updateSqlGenerate(target, where.Sql(), sets)
	stm, _ := db.Prepare(generate)
	defer func() {
		stm.Close()
	}()
	result, err := stm.Exec(values...)
	if err != nil {
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return count, err
	}
	return count, nil
}

// InsertModels
// 查询列表
// 参数说明:
//	target:条件载体，用户反射/填充等操作
//	where:where条件生成器,用于生成where条件
//	sets:需要修改的字段,必传
// 返回说明:
//	int64:插入生效的条目
//	error:异常信息
func InsertModels(target ...interface{}) (int64, error) {
	defer deferError("insert model method error")
	sql := insertSqlGenerate(target...)
	exec, err := db.Exec(sql)
	if err != nil {
		return 0, err
	}
	count, err := exec.RowsAffected()
	if err != nil {
		return count, err
	}
	return count, nil
}
