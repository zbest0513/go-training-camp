package main

import (
	"chapter05/cb"
	"fmt"
	"log"
)

func main() {

	//创建断路器，每20s 最大100个请求
	breaker := cb.CreateCircuitBreaker(100, 20)
	//创建执行器
	ds := new(cb.LocalExecutor)

	//构建请求的client
	client := cb.NewRemoteClient(breaker, ds)

	req := new(cb.Request)
	param := make(map[string]string, 5)
	param["code"] = "0"
	param["reason"] = "xxx"
	//param["sleep"] = "5"
	req.Param = param

	result := client.Execute(*req)
	log.Println(fmt.Sprintf("返回结果:{%v}", result))

}
