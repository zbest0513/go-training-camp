package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	//WaitGroupVersion() //wait group 实现的版本
	ErrGroupVersion() //errgroup 实现的版本
}

func signalHandle(f func()) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case msg := <-signals:
		log.Println(fmt.Sprintf("接受信号:%v", msg))
		f()
	}
}

//WaitGroupVersion
//WaitGroup版本并未考虑启动异常的情况和注销方法超时的情况
func WaitGroupVersion() {
	var wg sync.WaitGroup
	ctxA, cancelA := context.WithCancel(context.Background())
	defer cancelA()
	wg.Add(1)
	go server(ctxA, &wg, ":8001")
	wg.Add(1)
	go server(ctxA, &wg, ":8082")
	go signalHandle(func() {
		cancelA()
	})
	wg.Wait()
}

func server(ctx context.Context, wg *sync.WaitGroup, port string) {
	defer wg.Done()
	go http.ListenAndServe(port, nil)
	log.Println(fmt.Sprintf("listen port :%s", port))
	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("listen port :%s shutdown ......", port))
	}
}

//ErrGroupVersion
//ErrGroupVersion版本在WaitGroup版本基础上增加了
//1.server启动出错时退出所有
//2.shutdown程序的超时处理
func ErrGroupVersion() {
	ctx, ctxCancel := context.WithCancel(context.Background())
	g, ctxG := errgroup.WithContext(ctx)
	g.Go(func() error {
		return server2(ctxG, ":8081", ctxCancel)
	})
	g.Go(func() error {
		return server2(ctxG, ":8082", ctxCancel)
	})

	//监控kill信号
	//这里没有用Errgroup是因为上面server2处理的信号与signalHandle不同
	//避免server2收到信号退出而signalHandle 没办法退出导致g.Wait方法一直阻塞主线程
	go signalHandle(func() {
		ctxCancel()
	})
	g.Wait()
}

func server2(ctx context.Context, port string, cancelFunc context.CancelFunc) error {
	var err = errors.New("listen server run time exception")
	log.Println(fmt.Sprintf("listen port :%s", port))
	go func() {
		err2 := http.ListenAndServe(port, nil)
		if err2 != nil {
			log.Printf(fmt.Sprintf("http err :%v", err2))
			cancelFunc()
		}
	}()

	select {
	case <-ctx.Done():
		timeout, c := context.WithTimeout(context.TODO(), time.Second*time.Duration(5))
		go shutdown(port, c)
		//处理超时退出
		select {
		case <-timeout.Done():
			log.Println(fmt.Sprintf("listen port :%s shutdown finished", port))
		}
	}
	return err
}

//模拟注销时长
func shutdown(port string, cancelFunc context.CancelFunc) {
	rand.Seed(time.Now().UnixNano())
	t := rand.Intn(10)
	time.Sleep(time.Second * time.Duration(t))
	log.Println(fmt.Sprintf("listen port :%s shutdown use time %v second ......", port, t))
	cancelFunc()
}
