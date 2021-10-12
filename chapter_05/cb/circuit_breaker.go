package cb

import (
	"errors"
	"sync/atomic"
	"time"
)

type CircuitBreaker struct {
	isDestroy         bool      //是否销毁
	status            byte      //断路器状态:0-close/1-open/2-half open
	max               uint32    //最大并发数
	windowTime        int       //窗口时长，单位秒
	totalBucket       int       //总的桶数量
	currBucket        int       //当前桶位置
	buckets           []*Bucket //桶列表
	recoveryTime      int       //断路后尝试恢复时长，单位秒
	tryCloseTime      time.Time //尝试恢复时间
	minFailCount      int       //最小错误数
	failRate          float64   //错误率
	recoveryRate      float64   //恢复率阈值，当试错的
	tryTotal          uint32    //恢复期尝试总数
	leftover          *uint32   //剩余并发数
	totalSuccessCount *uint32   //时间窗口内总成功次数
	totalFailCount    *uint32   //时间窗口内总的失败次数
	totalCount        *uint32   //时间窗口内总的请求数量
	tryLeftover       *uint32   //恢复期尝试余数
	trySuccess        *uint32   //恢复期尝试成功数量
}

// Bucket
// leftover用于限制桶内流
// total、success、fail用于桶销毁后重置断路器内计数用的，扣减对应计数
type Bucket struct {
	max      uint32  //最大并发数
	total    *uint32 //总访问次数
	success  *uint32 //成功次数
	fail     *uint32 //失败次数
	leftover *uint32 //剩余量
}

func NewBucket(leftover uint32) *Bucket {
	var max = leftover
	var total = uint32(0)
	var success = uint32(0)
	var fail = uint32(0)
	return &Bucket{
		leftover: &max,
		total:    &total,
		success:  &success,
		fail:     &fail,
		max:      leftover,
	}
}

// CreateCircuitBreaker 创建默认断路器
func CreateCircuitBreaker(max uint32, windowTime int) *CircuitBreaker {
	defaultStrategy := new(DefaultStrategy)
	return NewCircuitBreaker(max, windowTime, defaultStrategy)
}

// NewCircuitBreaker 创建断路器
func NewCircuitBreaker(max uint32, windowTime int, strategy Strategy) *CircuitBreaker {
	//默认逻辑，1秒10个桶，每个桶处理100ms内的请求
	totalBucket := windowTime * 10
	buckets := make([]*Bucket, totalBucket, totalBucket)

	//计算每个桶的最大容量,取整
	bMax := max / uint32(totalBucket)
	for i := 0; i < totalBucket; i++ {
		buckets[i] = NewBucket(bMax)
	}

	var leftover uint32 = 0
	var totalSuccessCount uint32 = 0
	var totalCount uint32 = 0
	var totalFailCount uint32 = 0
	var tryLeftover = strategy.GetTryCount()
	var trySuccess uint32 = 0
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
	go func() {
		for !cb.isDestroy {
			<-t
			go func() {
				cb.nextWindow()
			}()
		}
	}()
}

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

func (cb *CircuitBreaker) nextWindow() {
	//将要抛弃的桶的计数归还给断路器
	bucket := cb.buckets[cb.currBucket]
	total := atomic.LoadUint32(bucket.total)
	atomic.AddUint32(cb.totalCount, ^uint32(-int32(total)-1))
	success := int32(atomic.LoadUint32(bucket.success))
	atomic.AddUint32(cb.totalSuccessCount, ^uint32(-success-1))
	fail := int32(atomic.LoadUint32(bucket.fail))
	atomic.AddUint32(cb.totalFailCount, ^uint32(-fail-1))
	atomic.AddUint32(cb.leftover, total)
	//重置桶
	cb.buckets[cb.currBucket] = NewBucket(bucket.max)
}

//断路器状态校验
func (cb *CircuitBreaker) check() {
	if cb.status == 1 { //开启状态
		//判断是否需要变更到半开状态
		if time.Now().Nanosecond() >= cb.tryCloseTime.Nanosecond() {
			cb.halfOpen()
		}
	} else if cb.status == 2 { //半开状态
		leftover := atomic.LoadUint32(cb.tryLeftover)
		success := float64(atomic.LoadUint32(cb.trySuccess))
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
		var failTotal = atomic.LoadUint32(cb.totalFailCount)
		if failTotal < uint32(cb.minFailCount) { //错误数量太少，不开启断路器
			return
		}
		var fails = float64(failTotal)
		var total = float64(atomic.LoadUint32(cb.totalCount))
		if fails/total < cb.failRate/100 { //错误率低，不开启断路器
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
	var one int32 = 1
	//总余量-1
	count := atomic.AddUint32(cb.leftover, ^uint32(-one-1))
	if count < 0 {
		//放行失败，余量补1
		atomic.AddUint32(cb.leftover, uint32(1))
		return 0, errors.New("触发总限额")
	}
	currBucket := cb.currBucket
	bucket := cb.buckets[currBucket]
	//桶余量-1
	count = atomic.AddUint32(bucket.leftover, ^uint32(-one-1))
	if count < 0 {
		//放行失败，余量补1
		atomic.AddUint32(cb.leftover, uint32(1))
		atomic.AddUint32(bucket.leftover, uint32(1))
		return 0, errors.New("触发桶限额")
	}

	//总放行+1
	atomic.AddUint32(cb.totalCount, uint32(1))
	//桶放行+1
	atomic.AddUint32(bucket.total, uint32(1))
	return currBucket, nil
}

// TryRecovery 试恢复
func (cb *CircuitBreaker) TryRecovery() (int, error) {
	var one int32 = 1
	//总余量-1
	count := atomic.AddUint32(cb.tryLeftover, ^uint32(-one-1))
	if count < 0 {
		//放行失败，余量补1
		atomic.AddUint32(cb.tryLeftover, uint32(1))
		return 0, errors.New("试恢复限额")
	}
	idx, err := cb.TryPass()
	if err != nil {
		//放行失败，余量补1
		atomic.AddUint32(cb.tryLeftover, uint32(1))
		return 0, err
	}
	return idx, nil
}

// SyncCounters 调用结果同步到断路器的计数
func (cb *CircuitBreaker) SyncCounters(isSuccess bool, idx int, flag bool) {
	atomic.AddUint32(cb.tryLeftover, uint32(1))
	if isSuccess {
		atomic.AddUint32(cb.totalSuccessCount, uint32(1))
		if atomic.LoadUint32(cb.buckets[idx].total) != 0 { //桶未过期，需要记录桶数据
			atomic.AddUint32(cb.buckets[idx].success, uint32(1))
			atomic.AddUint32(cb.buckets[idx].leftover, uint32(1))
		}
	} else {
		atomic.AddUint32(cb.totalFailCount, uint32(1))
		if atomic.LoadUint32(cb.buckets[idx].total) != 0 { //桶未过期，需要记录桶数据
			atomic.AddUint32(cb.buckets[idx].fail, uint32(1))
			atomic.AddUint32(cb.buckets[idx].leftover, uint32(1))
		}
	}
	if flag {
		cb.syncTryCounters(isSuccess)
	}
}

// 调用结果更新到恢复重试的计数上
func (cb *CircuitBreaker) syncTryCounters(isSuccess bool) {
	if isSuccess {
		atomic.AddUint32(cb.trySuccess, uint32(1))
	}
}
