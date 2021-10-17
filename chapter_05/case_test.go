package main

import (
	"chapter05/cb"
	"fmt"
	"log"
	"sync/atomic"
	"testing"
	"time"
)

//每135ms 执行一次f函数，执行n次
func scheduler(n int, f func()) {
	t := time.Tick(135 * time.Millisecond)
	//100ms 生成一个新桶
	i := 0
	for i < n {
		i++
		go func() {
			<-t
			f()
		}()
	}
}

//等待n秒 ，防止主线程退出
func waitSecond(n int) {
	t := time.Tick(time.Duration(n) * time.Second)
	flag := true
	for flag {
		flag = false
		<-t
		go func() {
			println("结束等待程序")
		}()
	}
}

//全部成功的数据初始化
func successDataSetup(count int, t string) []cb.Request {
	requests := make([]cb.Request, count, count)
	for i := 0; i < count; i++ {
		req := new(cb.Request)
		param := make(map[string]string, 5)
		param["code"] = "0"
		param["reason"] = "xxx"
		if "" != t {
			param["sleep"] = t
		}
		req.Param = param
		requests[i] = *req
	}
	return requests
}

func failDataSetup(count int, t string) []cb.Request {
	requests := make([]cb.Request, count, count)
	for i := 0; i < count; i++ {
		req := new(cb.Request)
		param := make(map[string]string, 5)
		param["code"] = "1"
		param["reason"] = "xxx"
		if "" != t {
			param["sleep"] = t
		}
		req.Param = param
		requests[i] = *req
	}
	return requests
}

//test case: 测试全放行
// 20s的窗口最大最大并行 : 100000
// 100ms的桶窗口最大并行 : 500 （100000/20/10=500）
// 初始化200个请求 并发执行 预计会在第一个桶内都受理
func TestCurrent(t *testing.T) {
	//创建断路器，每20s 最大100000个请求
	breaker := cb.CreateCircuitBreaker(100000, 20)
	//创建执行器
	ds := new(cb.LocalExecutor)

	//构建请求的client
	client := cb.NewRemoteClient(breaker, ds)

	//初始化200个成功的请求
	requests := successDataSetup(200, "")

	for _, request := range requests {
		go func() {
			_ = client.Execute(request)
		}()
	}

	waitSecond(1)
}

//test case: 测试桶限流
// 20s的窗口最大最大并行 : 100000
// 100ms的桶窗口最大并行 : 500 （100000/20/10=500）
// 初始化505个请求 并发执行 预计会在第一个桶内都受理500个，降级返回5个
func TestBucketLimit(t *testing.T) {
	//创建断路器，每20s 最大100000个请求
	breaker := cb.CreateCircuitBreaker(100000, 20)
	//创建执行器
	ds := new(cb.LocalExecutor)

	//构建请求的client
	client := cb.NewRemoteClient(breaker, ds)

	//初始化200个成功的请求
	requests := successDataSetup(505, "")

	list := make([]cb.Result, 5, 5)

	var idx int32 = 0

	for _, request := range requests {
		go func() {
			result := client.Execute(request)
			if result.Code == 5 { //降级的返回
				idxA := atomic.AddInt32(&idx, int32(1))
				if idxA < int32(len(list)+1) {
					list[int(idxA)-1] = result
				}
			}
		}()
	}

	waitSecond(1)
	log.Println(fmt.Sprintf("降级的返回:[%v]", list))
}

//test case: 测试总限流
// 4s的窗口最大最大并行 : 1000
// 100ms的桶窗口最大并行 : 25 （1000/4/10 = 25）
// 每100ms 发送 25个请求 并发执行 每个请求睡5s，第5秒应该触发总限流
func TestTotalLimit(t *testing.T) {

	//创建断路器，每4s 最大1000个请求
	breaker := cb.CreateCircuitBreaker(1000, 4)
	//创建执行器
	ds := new(cb.LocalExecutor)
	//构建请求的client
	client := cb.NewRemoteClient(breaker, ds)
	max := 100
	list := make([]cb.Result, max, max)
	var idx int32 = 0
	//初始化25个成功的请求
	requests := successDataSetup(25, "5")
	//每110ms 发起25个请求 ，执行5s
	//用110ms做步长避免和桶生成的110秒想撞，造成桶限流
	scheduler(28*10, func() {
		for _, request := range requests {
			go func() {
				result := client.Execute(request)
				if result.Code == 5 { //降级的返回
					idxA := atomic.AddInt32(&idx, int32(1))
					if idxA < int32(len(list)+1) {
						list[int(idxA)-1] = result
					}
				}
			}()
		}
	})
	waitSecond(30)
	log.Println(fmt.Sprintf("降级的返回:[%v]", idx))
}

//test case: 测试断路器自动开启
// 默认策略：错误次数超过20个，错误率达到30%
// 所以需要第一个桶 有30个请求，其中20个是错的，40个对的
// 100ms的桶窗口最大并行 : 60
// 4s的窗口最大最大并行 :  2500  （60*10*4 = 2400） 为了让桶的并发更高一点，2400 扩大到2500方便测试
// 直接发20个失败的 40个成功的 看看断路器是否开启
// 等待20s 看看会不会在开启5s（默认策略是5s）后变为半开状态
func TestFuse(t *testing.T) {

	//创建断路器，每4s 最大2500个请求
	breaker := cb.CreateCircuitBreaker(2500, 4)
	//创建执行器
	ds := new(cb.LocalExecutor)
	//构建请求的client
	client := cb.NewRemoteClient(breaker, ds)

	fails := failDataSetup(20, "")

	for _, request := range fails {
		go func() {
			_ = client.Execute(request)
		}()
	}
	successes := successDataSetup(40, "")
	for _, request := range successes {
		go func() {
			_ = client.Execute(request)
		}()
	}

	waitSecond(7)
}

//test case: 测试断路器开启后自动关闭
// 默认策略：错误次数超过20个，错误率达到30%
// 所以需要第一个桶 有30个请求，其中20个是错的，40个对的
// 100ms的桶窗口最大并行 : 60
// 4s的窗口最大最大并行 :  2500  （60*10*4 = 2400） 为了让桶的并发更高一点，2400 扩大到2500方便测试
// 直接发20个失败的 40个成功的 看看断路器开启
// 等待6（5秒就变为半开）秒，放5个成功的请求，测试时候会变为关闭
func TestAutoClose(t *testing.T) {

	//创建断路器，每4s 最大2500个请求
	breaker := cb.CreateCircuitBreaker(2500, 4)
	//创建执行器
	ds := new(cb.LocalExecutor)
	//构建请求的client
	client := cb.NewRemoteClient(breaker, ds)

	fails := failDataSetup(20, "")

	for _, request := range fails {
		go func() {
			_ = client.Execute(request)
		}()
	}
	successes := successDataSetup(40, "")
	for _, request := range successes {
		go func() {
			_ = client.Execute(request)
		}()
	}

	waitSecond(6)
	successes2 := successDataSetup(5, "")
	for _, request := range successes2 {
		go func() {
			_ = client.Execute(request)
		}()
	}
	waitSecond(20)
}
