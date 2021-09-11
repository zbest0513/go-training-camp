package main

import (
	"fmt"
	"log"
)

func main() {

	execute(useCase1)

}

func execute(useCase func()) {
	//每个请求都在入口把panic catch住
	//避免一个请求影响整个server进程
	defer catch()
	//执行函数
	useCase()
}

func useCase1() {
	log.Println("-----------")
}

func catch() {
	if r := recover(); r != nil {
		log.Println(fmt.Sprintf("%+v\n", r))
	}
}
