package commands

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var StartCommand = cli.Command{
	Name:   "start",
	Usage:  "start a web server",
	Action: start,
}

func start(c *cli.Context) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	go Server(ctx)
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT) // 这里注册一下结束信号处理管道
	<-quit                                              // 无缓冲channel，在没有接收到信号之前后面的操作都会被阻塞你
	cancel()
	time.Sleep(time.Second * 1)
	return
}

func Server(ctx context.Context) {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	serve := http.Server{
		Addr:    "127.0.0.1:1234",
		Handler: engine,
	}
	// 这里用goroutine的原因是因为：ListenAndServe是一个死循环会导致后面无法运行
	go func() {
		if err := serve.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				fmt.Println("web server shutdown complete")
			} else {
				fmt.Printf("web server closed unexpect: %s", err)
			}
		}
	}()
	<-ctx.Done()
	fmt.Println("sutdown web server")
	if err := serve.Close(); err != nil {
		fmt.Printf("web server shutdown failed: %v", err)
	}
}
