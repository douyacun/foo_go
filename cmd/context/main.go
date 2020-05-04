package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx)
	time.Sleep(time.Second * 10)
	fmt.Println("close web server")
	cancel()
}

func watch(ctx context.Context) {
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
	fmt.Printf("sutdown web server")
	if err := serve.Close(); err != nil {
		fmt.Printf("web server shutdown failed: %v", err)
	}
}

