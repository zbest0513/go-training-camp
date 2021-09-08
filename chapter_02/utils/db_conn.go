package utils

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	USERNAME        = "root"
	PASSWORD        = "root"
	NETWORK         = "tcp"
	SERVER          = "localhost"
	PORT            = 3306
	DATABASE        = "test"
	MAXOPENCONNS    = 100
	MAXIDLECONNS    = 16
	CONNMAXLIFETIME = 100 * time.Second
)

var db = &sql.DB{}
var err error

func init() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Open mysql failed,err:%v\n", err))
	}
	db.SetConnMaxLifetime(CONNMAXLIFETIME)
	db.SetMaxOpenConns(MAXOPENCONNS)
	db.SetMaxIdleConns(MAXIDLECONNS)
	fmt.Println("初始化数据库连接成功...")
}

func GetConn() *sql.DB {
	return db
}
