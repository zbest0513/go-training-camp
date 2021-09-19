package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func server(ctx context.Context, wg *sync.WaitGroup, port string) {
	defer wg.Done()
	go http.ListenAndServe(port, nil)
	log.Println(fmt.Sprintf("listen port :%s", port))
	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("listen port :%s shutdown ......", port))
	}
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

func main() {
	//WaitGroupVersion() //wait group 实现的版本
	ErrGroupVersion() //errgroup 实现的版本
}

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

func catchRuntimeException() {
	if r := recover(); r != nil {
		log.Println(fmt.Sprintf("catch runtime exception"))
	}
}

func server2(ctx context.Context, port string) error {
	//捕捉到运行时panic,忽略掉直接返回err，避免直接程序退出
	defer catchRuntimeException()
	var err = errors.New("listen server run time exception")
	g2, _ := errgroup.WithContext(ctx)
	log.Println(fmt.Sprintf("listen port :%s", port))
	g2.Go(func() error {
		err2 := http.ListenAndServe(port, nil)
		log.Println(fmt.Sprintf("http err :%v", err2))
		return err2
	})
	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("listen port :%s shutdown ......", port))
	}
	//err = g2.Wait()
	return err
}

func ErrGroupVersion() {
	ctx, ctxCancel := context.WithCancel(context.Background())
	g, ctxG := errgroup.WithContext(ctx)
	g.Go(func() error {
		return server2(ctxG, ":8081")
	})
	g.Go(func() error {
		return server2(ctxG, ":8082")
	})
	//监控kill信号
	go signalHandle(func() {
		ctxCancel()
	})

	g.Wait()
}
