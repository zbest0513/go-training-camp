package utils

import "fmt"

//链式操作，生成where 条件
//功能待完善
//目前只支持 where xxx = xxx and xx = xx

type WhereGenerator struct {
	whereSql string
}

func (receiver *WhereGenerator) NewInstance() *WhereGenerator {
	receiver.whereSql = " where 1 = 1 "
	return receiver
}

func (receiver *WhereGenerator) And(column string) *WhereGenerator {
	receiver.whereSql = fmt.Sprintf("%s and %s ", receiver.whereSql, column)
	return receiver
}

func (receiver *WhereGenerator) Equals(value interface{}) *WhereGenerator {
	receiver.whereSql = fmt.Sprintf("%s = '%s' ", receiver.whereSql, value)
	return receiver
}

func (receiver *WhereGenerator) Sql() string {
	return receiver.whereSql
}
