package cb

import (
	"fmt"
	"log"
)

type RemoteClient struct {
	cb       *CircuitBreaker //断路器
	executor Executor        //请求执行器
}

func (client *RemoteClient) Execute(req Request) Result {
	if client.cb == nil { //没有断路器
		return client.executor.execute(req)
	}
	status := client.cb.GetStatus()

	var flag = true //false为试错流量

	if status == 1 { //断路器开启，走降级
		log.Println("断路器开启流量被熔断")
		return req.Fallback.F()
	} else if status == 2 { //半开
		flag = false
	}
	//流量计数
	idx, err := client.cb.Pass(flag)
	if err != nil {
		log.Println(fmt.Sprintf("流量被熔断:%v", err))
		return req.Fallback.F()
	}
	//执行调用
	result := client.executor.execute(req)
	//将结果同步给断路器计数
	client.cb.SyncCounters(result.Code == 0, idx, flag)
	return result
}

func NewRemoteClient(cb *CircuitBreaker, executor Executor) *RemoteClient {
	return &RemoteClient{
		cb:       cb,
		executor: executor,
	}
}
