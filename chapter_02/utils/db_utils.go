package utils

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
)

type DBUtils struct {
	isTrans bool    //是否开启事务
	tx      *sql.Tx //事务管理器
	msg     error   //错误信息
}

//QueryOne
// 查询单条记录
// 参数说明:
//	target:条件载体，用户反射/填充等操作
//	where:where条件生成器,用于生成where条件
//	scan:需要检索的字段
// 返回说明:
//	interface:检索到的条目
//	error:异常信息
func (receiver *DBUtils) QueryOne(target interface{}, where *WhereGenerator, scans ...string) (interface{}, error) {
	return queryOne(target, where, scans...)
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
func (receiver *DBUtils) QueryList(target interface{}, where *WhereGenerator, scans ...string) ([]interface{}, error, int) {
	return queryList(target, where, scans...)
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
func (receiver *DBUtils) UpdateModels(target interface{}, where *WhereGenerator, sets []string) (int64, error) {
	return updateModels(receiver, target, where, sets)
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
func (receiver *DBUtils) InsertModels(target ...interface{}) (int64, error) {
	return insertModels(receiver, target...)
}

// DeleteModels
// 删除
// 参数说明:
//	target:条件载体，用户反射/填充等操作
//	where:where条件生成器,用于生成where条件
// 返回说明:
//	int64:更新生效的条目数量
//	error:异常信息
func (receiver *DBUtils) DeleteModels(target interface{}, where *WhereGenerator) (int64, error) {
	return deleteModels(receiver, target, where)
}

//TxExec
// 支持事务执行
// method 可传入update、insert、delete函数
func (receiver *DBUtils) TxExec(executors ...*TransTaskExecutor) (int64, error) {
	receiver.isTrans = true
	receiver.tx, receiver.msg = GetConn().Begin()
	var result int64 = 0
	for _, execute := range executors {
		count, err := execute.exec(receiver)
		if err != nil {
			receiver.tx.Rollback()
			return count, err
		}
		result += count
	}
	err := receiver.tx.Commit()
	if err != nil {
		receiver.tx.Rollback()
		return result, err
	}
	return result, nil
}

func queryOne(target interface{}, where *WhereGenerator, scans ...string) (interface{}, error) {
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

func queryList(target interface{}, where *WhereGenerator, scans ...string) ([]interface{}, error, int) {
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

func updateModels(dbUtils *DBUtils, target interface{}, where *WhereGenerator, sets []string) (int64, error) {
	defer deferError("update model method error")
	generate, values := updateSqlGenerate(target, where.Sql(), sets)

	var stm *sql.Stmt
	var err error

	if dbUtils.isTrans {
		stm, err = dbUtils.tx.Prepare(generate)
	} else {
		stm, err = db.Prepare(generate)
	}

	if err != nil {
		return 0, err
	}
	defer func() {
		err := stm.Close()
		log.Println(fmt.Sprintf("stm close err :%+v", err))
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

func insertModels(dbUtils *DBUtils, target ...interface{}) (int64, error) {
	defer deferError("insert model method error")
	generate := insertSqlGenerate(target...)
	var exec sql.Result
	var err error
	if dbUtils.isTrans {
		exec, err = dbUtils.tx.Exec(generate)
	} else {
		exec, err = db.Exec(generate)
	}
	if err != nil {
		return 0, err
	}
	count, err := exec.RowsAffected()
	if err != nil {
		return count, err
	}
	return count, nil
}

func deleteModels(dbUtils *DBUtils, target interface{}, where *WhereGenerator) (int64, error) {
	defer deferError("delete model method error")
	generate := deleteSqlGenerate(target, where.Sql())

	var result sql.Result
	var err error
	if dbUtils.isTrans {
		result, err = dbUtils.tx.Exec(generate)
	} else {
		result, err = db.Exec(generate)
	}
	if err != nil {
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return count, err
	}
	return count, nil
}

type TransTaskExecutor struct {
	method func(...interface{}) (int64, error)
	args   []interface{}
}

func (receiver *TransTaskExecutor) exec(db *DBUtils) (int64, error) {
	return receiver.method(db, receiver.args)
}

func (receiver *TransTaskExecutor) NewUpdateTaskExecutor(method func(...interface{}) (int64, error), target interface{}, where *WhereGenerator, sets []string) *TransTaskExecutor {
	args := make([]interface{}, 3, 3)
	args[0] = true
	args[2] = target
	args[3] = where
	args[4] = sets

	return &TransTaskExecutor{
		args:   args,
		method: method,
	}
}

func (receiver *TransTaskExecutor) NewInsertTaskExecutor(method func(...interface{}) (int64, error), target ...interface{}) *TransTaskExecutor {
	count := len(target)
	args := make([]interface{}, count, count)
	for i := 0; i < count; i++ {
		args[i] = target[i]
	}
	return &TransTaskExecutor{
		args:   args,
		method: method,
	}
}

func (receiver *TransTaskExecutor) NewDeleteTaskExecutor(method func(...interface{}) (int64, error), target interface{}, where *WhereGenerator) *TransTaskExecutor {

	args := make([]interface{}, 4, 4)
	args[0] = true
	args[2] = target
	args[3] = where

	return &TransTaskExecutor{
		args:   args,
		method: method,
	}
}
