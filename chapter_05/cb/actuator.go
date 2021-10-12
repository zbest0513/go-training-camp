package cb

import (
	"strconv"
	"time"
)

type Executor interface {
	execute(req Request) Result
}

// HttpExecutor http 请求器
type HttpExecutor struct {
	//TODO 集成 resty或其他httpclient框架
}

func (receiver *HttpExecutor) execute(req Request) Result {
	return Result{}
}

// GrpcExecutor gprc 请求器
type GrpcExecutor struct {
	//TODO 集成 gprc框架
}

func (receiver *GrpcExecutor) execute(req Request) Result {
	return Result{}
}

// LocalExecutor 不调用远程，用于测试断路器的
// req Request 的param 接收一个map[string]string
// 直接将req 转化成 Result 测试断路器
type LocalExecutor struct {
}

func (receiver *LocalExecutor) execute(req Request) Result {
	//将参数解析成result 返回
	param := req.Param
	m, ok := param.(map[string]string)
	if !ok {
		return Result{1, "非法参数", nil}
	}
	code, err := strconv.Atoi(m["code"])
	if err != nil {
		return Result{2, "非法参数", nil}
	}
	sleep, err := strconv.Atoi(m["sleep"])
	if err != nil {
		//跳过
	} else {
		//测试超时
		time.Sleep(time.Duration(sleep) * time.Second)
	}
	return Result{
		Code:   code,
		Reason: m["reason"],
		Data:   nil,
	}
}
