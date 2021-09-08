package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	username        = "root"
	password        = "zhangbin123"
	network         = "tcp"
	server          = "mysql.zbest.tech"
	port            = 13306
	database        = "notify"
	maxopenconns    = 100
	maxidleconns    = 16
	connmaxlifetime = 100 * time.Second
)

var db = &sql.DB{}
var err error

func init() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", username, password, network, server, port, database)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Open mysql failed,err:%v\n", err))
	}
	db.SetConnMaxLifetime(connmaxlifetime)
	db.SetMaxOpenConns(maxopenconns)
	db.SetMaxIdleConns(maxidleconns)
	fmt.Println("初始化数据库连接成功...")
}

func GetConn() *sql.DB {
	return db
}
