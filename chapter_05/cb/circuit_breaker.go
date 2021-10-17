package cb

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type CircuitBreaker struct {
	isDestroy         bool      //是否销毁
	status            byte      //断路器状态:0-close/1-open/2-half open
	max               int64     //最大并发数
	windowTime        int       //窗口时长，单位秒
	totalBucket       int       //总的桶数量
	currBucket        int       //当前桶位置
	buckets           []*Bucket //桶列表
	recoveryTime      int       //断路后尝试恢复时长，单位秒
	tryCloseTime      time.Time //尝试恢复时间
	minFailCount      int       //最小错误数
	failRate          float64   //错误率
	recoveryRate      float64   //恢复率阈值，当试错的
	tryTotal          int64     //恢复期尝试总数
	leftover          *int64    //剩余并发数
	totalSuccessCount *int64    //时间窗口内总成功次数
	totalFailCount    *int64    //时间窗口内总的失败次数
	totalCount        *int64    //时间窗口内总的请求数量
	tryLeftover       *int64    //恢复期尝试余数
	trySuccess        *int64    //恢复期尝试成功数量
}

// Bucket
// leftover用于限制桶内流
// total、success、fail用于桶销毁后重置断路器内计数用的，扣减对应计数
type Bucket struct {
	max      int64  //最大并发数
	total    *int64 //总访问次数
	success  *int64 //成功次数
	fail     *int64 //失败次数
	leftover *int64 //剩余量
}

func NewBucket(leftover int64) *Bucket {
	var max = leftover
	var total = int64(0)
	var success = int64(0)
	var fail = int64(0)
	return &Bucket{
		leftover: &max,
		total:    &total,
		success:  &success,
		fail:     &fail,
		max:      leftover,
	}
}

// CreateCircuitBreaker 创建默认断路器
func CreateCircuitBreaker(max int64, windowTime int) *CircuitBreaker {
	defaultStrategy := new(DefaultStrategy)
	return NewCircuitBreaker(max, windowTime, defaultStrategy)
}

// NewCircuitBreaker 创建断路器
func NewCircuitBreaker(max int64, windowTime int, strategy Strategy) *CircuitBreaker {
	//默认逻辑，1秒10个桶，每个桶处理100ms内的请求
	totalBucket := windowTime * 10
	buckets := make([]*Bucket, totalBucket, totalBucket)

	//计算每个桶的最大容量,取整
	bMax := max / int64(totalBucket)
	for i := 0; i < totalBucket; i++ {
		buckets[i] = NewBucket(bMax)
	}

	var leftover = max
	var totalSuccessCount int64 = 0
	var totalCount int64 = 0
	var totalFailCount int64 = 0
	var tryLeftover = strategy.GetTryCount()
	var trySuccess int64 = 0
	cb := &CircuitBreaker{
		status:            0,
		max:               max,
		windowTime:        windowTime,
		leftover:          &leftover,
		totalBucket:       totalBucket,
		currBucket:        0,
		buckets:           buckets,
		recoveryTime:      strategy.GetRecoveryTime(),
		minFailCount:      strategy.GetMinFailCount(),
		failRate:          strategy.GetFailRate(),
		tryTotal:          strategy.GetTryCount(),
		recoveryRate:      strategy.GetRecoveryRate(),
		totalSuccessCount: &totalSuccessCount,
		totalCount:        &totalCount,
		totalFailCount:    &totalFailCount,
		tryLeftover:       &tryLeftover,
		trySuccess:        &trySuccess,
		isDestroy:         false,
	}
	cb.start()
	return cb
}

func (cb *CircuitBreaker) start() {
	t := time.Tick(100 * time.Millisecond)
	//100ms 生成一个新桶
	go func() {
		for !cb.isDestroy {
			<-t
			go func() {
				cb.nextWindow()
			}()
		}
	}()
}

// Destroy 停止生成新桶
func (cb *CircuitBreaker) Destroy() {
	cb.isDestroy = true
}

func (cb *CircuitBreaker) open() {
	cb.status = 1
	var t = cb.tryTotal
	cb.tryLeftover = &t //重置试错余量
	//重置下次尝试恢复时间
	cb.tryCloseTime = time.Now().Add(time.Second * time.Duration(cb.recoveryTime))
}

func (cb *CircuitBreaker) close() {
	cb.status = 0
}

func (cb *CircuitBreaker) halfOpen() {
	cb.status = 2
}

// 生成新桶，丢弃老桶
func (cb *CircuitBreaker) nextWindow() {
	//将要抛弃的桶的计数归还给断路器
	currBucket := cb.currBucket
	bucket := cb.buckets[currBucket]
	//将丢弃的桶内，总次数、成功和失败的次数，从断路器的计数器中扣除
	total := load(bucket.total)
	success := load(bucket.success)
	less(cb.totalSuccessCount, success)
	fail := load(bucket.fail)
	less(cb.totalFailCount, fail)
	less(cb.totalCount, success)
	less(cb.totalCount, fail)
	log.Println(fmt.Sprintf("桶:%v[受理请求:%v,桶余量:%v,总余量:%v],时间窗口:[总受理:%v,成功：%v,失败:%v,状态:%v]",
		currBucket, total, load(bucket.leftover), load(cb.leftover), load(cb.totalCount), load(cb.totalSuccessCount),
		load(cb.totalFailCount), cb.status))
	//重置桶
	cb.buckets[cb.currBucket] = NewBucket(bucket.max)
	cb.currBucket = (cb.currBucket + 1) % cb.totalBucket

	cb.check()
}

//断路器状态校验
func (cb *CircuitBreaker) check() {
	if cb.status == 1 { //开启状态
		//判断是否需要变更到半开状态
		if time.Now().UnixNano() >= cb.tryCloseTime.UnixNano() {
			cb.halfOpen()
		}
	} else if cb.status == 2 { //半开状态
		leftover := load(cb.tryLeftover)
		success := float64(load(cb.trySuccess))
		total := float64(cb.tryTotal)
		//全部试完
		if leftover == 0 {
			if success/total >= cb.recoveryRate/100 { //关闭
				cb.close()
			} else { //再次开启
				cb.open()
			}
		}
	} else { //关闭状态
		var failTotal = load(cb.totalFailCount)
		if failTotal < int64(cb.minFailCount) { //错误数量太少，不开启断路器
			return
		}
		var fails = float64(failTotal)
		var success = float64(load(cb.totalSuccessCount))
		if fails/(success+fails) < cb.failRate/100 { //错误率低，不开启断路器
			return
		}
		cb.open() //开启断路器
	}
}

func (cb *CircuitBreaker) GetStatus() byte {
	return cb.status
}

func (cb *CircuitBreaker) Pass(flag bool) (int, error) {
	if flag {
		return cb.TryPass()
	} else {
		return cb.TryRecovery()
	}
}

// TryPass 试放行
func (cb *CircuitBreaker) TryPass() (int, error) {
	//总余量-1
	count := less(cb.leftover)
	if count < 0 {
		//放行失败，余量补1
		add(cb.leftover)
		return 0, errors.New("触发总限额")
	}
	currBucket := cb.currBucket
	bucket := cb.buckets[currBucket]
	//桶余量-1
	count = less(bucket.leftover)
	if count < 0 {
		//放行失败，余量补1
		add(cb.leftover)
		add(bucket.leftover)
		return 0, errors.New("触发桶限额")
	}

	//总放行+1
	add(cb.totalCount)
	//桶放行+1
	add(bucket.total)
	return currBucket, nil
}

// TryRecovery 试恢复
func (cb *CircuitBreaker) TryRecovery() (int, error) {
	//总余量-1
	count := less(cb.tryLeftover)
	if count < 0 {
		//放行失败，余量补1
		add(cb.tryLeftover)
		return 0, errors.New("试恢复限额")
	}
	idx, err := cb.TryPass()
	if err != nil {
		//放行失败，余量补1
		add(cb.tryLeftover)
		return 0, err
	}
	return idx, nil
}

// SyncCounters 调用结果同步到断路器的计数
func (cb *CircuitBreaker) SyncCounters(isSuccess bool, idx int, flag bool) {
	add(cb.leftover) //有返回结果 + 总的余量
	//桶的余量增加要限制，可能上一轮桶的请求
	count := add(cb.buckets[idx].leftover)
	if count > cb.buckets[idx].max { //加多了 减回来
		less(cb.buckets[idx].leftover)
	}
	if isSuccess { //成功的记录 桶的成功
		add(cb.totalSuccessCount)
		add(cb.buckets[idx].success)
	} else {
		add(cb.totalFailCount)
		add(cb.buckets[idx].fail)
	}
	if flag {
		cb.syncTryCounters(isSuccess)
	}
	//回写完结果判断状态
	cb.check()
}

// 调用结果更新到恢复重试的计数上
func (cb *CircuitBreaker) syncTryCounters(isSuccess bool) {
	if isSuccess {
		add(cb.trySuccess)
	}
}
