package main

import (
	"chapter_02/dao"
	"log"
)

func main() {
	isExist, err := dao.CheckUser("张三")
	if err != nil {
		log.Println("请求异常")
		return
	}
	if isExist {
		log.Println("用户张三存在")
	} else {
		log.Println("用户张三不存在")
	}
}
