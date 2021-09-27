package pkg

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"math/rand"
	"notify-server/internal/service/router"
	"notify/pkg/config"
	gutils "notify/pkg/utils"
	"time"
)

var ProviderSetPkg = wire.NewSet(NewApp)

type App struct {
	config        *config.Config
	MonitorRouter *router.MonitorRouter
	ManageRouter  *router.ManageRouter
}

func NewApp(config *config.Config, MonitorRouter *router.MonitorRouter, ManageRouter *router.ManageRouter) (*App, error) {
	return &App{
		config:        config,
		MonitorRouter: MonitorRouter,
		ManageRouter:  ManageRouter,
	}, nil
}

func (app *App) Run() error {
	ctx, ctxCancel := context.WithCancel(context.Background())
	g, ctxG := errgroup.WithContext(ctx)
	vip := app.config.Vip
	g.Go(func() error {
		port := vip.GetInt("servers.manage-server.port")
		engine := gin.New()
		app.ManageRouter.Register(engine)
		return app.server(ctxG, engine, fmt.Sprintf(":%v", port), ctxCancel)
	})
	g.Go(func() error {
		port := vip.GetInt("servers.monitor-server.port")
		engine := gin.New()
		app.MonitorRouter.Register(engine)
		return app.server(ctxG, engine, fmt.Sprintf(":%v", port), ctxCancel)
	})
	//监控kill信号
	//这里没有用Errgroup是因为上面server2处理的信号与signalHandle不同
	//避免server2收到信号退出而signalHandle 没办法退出导致g.Wait方法一直阻塞主线程
	go gutils.SignalHandle(func() {
		ctxCancel()
	})
	return g.Wait()
}

func (app *App) server(ctx context.Context, engine *gin.Engine, port string, cancelFunc context.CancelFunc) error {
	var err = errors.New("listen server run time exception")
	vip := app.config.Vip
	shutdownTimeout := vip.GetInt("servers.shutdown-timeout")
	log.Println(fmt.Sprintf("listen port :%s", port))
	go func() {
		if err2 := engine.Run(port); err2 != nil {
			log.Printf(fmt.Sprintf("http err :%v", err2))
			cancelFunc()
		}
	}()
	select {
	case <-ctx.Done():
		timeout, c := context.WithTimeout(context.TODO(), time.Second*time.Duration(shutdownTimeout))
		go shutdown(port, c)
		//处理超时退出
		select {
		case <-timeout.Done():
			log.Println(fmt.Sprintf("listen port :%s shutdown finished", port))
			err = nil
		}
	}
	return err
}

func shutdown(port string, cancelFunc context.CancelFunc) {
	rand.Seed(time.Now().UnixNano())
	//t := rand.Intn(10)
	//time.Sleep(time.Second * time.Duration(t))
	log.Println(fmt.Sprintf("listen port :%s shutdown ......", port))
	cancelFunc()
}
