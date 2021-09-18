package demo

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"log"
	"sync"
	"testing"
	"time"
)

func sleepFunc(long int) {
	time.Sleep(time.Second * time.Duration(long))
	log.Println("sleepFunc .......")
}

func exceptionFunc(msg string) {
	panic(msg)
}

func fastFunc() {
	log.Println("fastFunc .......")
}

func TestWaitGroup(t *testing.T) {

	defer func() {
		log.Println("==========")
		if r := recover(); r != nil {
			log.Fatalf("======%v", r)
		}
	}()

	group := sync.WaitGroup{}
	group.Add(1)
	go func() {
		sleepFunc(20)
		group.Done()
	}()
	group.Add(1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("catch----------")
			}
		}()
		exceptionFunc("xxx")
		group.Done()
	}()

	group.Add(1)
	go func() {
		fastFunc()
		group.Done()
	}()

	group.Wait()

}

func errFunc() error {
	return errors.New("错误")
}

func handleErr(err error) {
	if r := recover(); r != nil {
		err = errors.New("出错了")
	}
}

func dowait20s(i int) int {
	time.Sleep(time.Second * 4)
	return i * 20
}

func TestChannel(t *testing.T) {
	ints := make(chan int)
	go func() {
		ints <- dowait20s(20)
	}()
	select {
	case x := <-ints:
		log.Printf("i : %v", x)
	case <-time.After(time.Second * 5):
		println("time out")
	}
}

func TestErrGroup(t *testing.T) {

	var g errgroup.Group

	g.Go(func() error {
		var err error
		defer handleErr(err)
		sleepFunc(20)
		return err
	})

	g.Go(func() error {
		//var err error
		//defer handleErr(err)
		return errFunc()
	})

	g.Go(func() error {
		var err error
		defer handleErr(err)
		fastFunc()
		return err
	})

	err := g.Wait()

	log.Printf("error .............%v", err)

}

func TestWithContext(t *testing.T) {

	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		var err error
		defer handleErr(err)
		sleepFunc(20)
		return err
	})

	g.Go(func() error {
		//var err error
		//defer handleErr(err)
		err := errFunc()
		if err != nil {
			ctx.Done()
		}
		return err
	})

	g.Go(func() error {
		var err error
		defer handleErr(err)
		fastFunc()
		return err
	})

	err := g.Wait()

	log.Printf("error .............%v", err)

}
