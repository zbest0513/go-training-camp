package cb

type Strategy interface {
	GetRecoveryTime() int     //N秒后尝试恢复
	GetMinFailCount() int     //最小失败次数
	GetFailRate() float64     //单位%
	GetTryCount() int64       //尝试恢复的请求数量
	GetRecoveryRate() float64 //尝试请求成功率超过N，会由半开恢复为关闭，单位%
}

type DefaultStrategy struct {
}

func (receiver DefaultStrategy) GetRecoveryTime() int {
	return 5
}

func (receiver DefaultStrategy) GetMinFailCount() int {
	return 20
}

func (receiver DefaultStrategy) GetFailRate() float64 {
	return 30
}

func (receiver DefaultStrategy) GetTryCount() int64 {
	return 5
}

func (receiver DefaultStrategy) GetRecoveryRate() float64 {
	return 80
}
