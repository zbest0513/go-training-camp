package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
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

func init() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", username, password, network, server, port, database)
	db, errMsg = sql.Open("mysql", dsn)
	if errMsg != nil {
		//重要资源无法获取连接，不可恢复异常
		log.Panic(fmt.Sprintf("open mysql failed,err:\n%+v\n", errMsg))
	}
	db.SetConnMaxLifetime(connmaxlifetime)
	db.SetMaxOpenConns(maxopenconns)
	db.SetMaxIdleConns(maxidleconns)
	log.Println("初始化数据库连接成功...")
}

func GetConn() *sql.DB {
	return db
}
