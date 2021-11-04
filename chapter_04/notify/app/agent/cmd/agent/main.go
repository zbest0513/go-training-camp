package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"notify-agent/internal"
	gutils "notify/pkg/utils"
)

func main() {

	ctx, ctxCancel := context.WithCancel(context.Background())
	g, _ := errgroup.WithContext(ctx)
	processor := internal.NewAgentProcessor()
	g.Go(func() error {
		processor.Start()
		return nil
	})
	go gutils.SignalHandle(func() {
		log.Println("优雅退出")
		processor.Destroy()
		ctxCancel()
	})
	g.Wait()
}
