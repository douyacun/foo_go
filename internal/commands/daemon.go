package commands

import (
	"context"
	"fmt"
	"github.com/sevlyar/go-daemon"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var DaemonCommend = cli.Command{
	Name:  "start",
	Usage: "detach from the console (daemon mode)",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "pid-filename,p",
			Value: "../runtime/run/foo.pid",
			Usage: "language for the greeting",
		},
		cli.StringFlag{
			Name:  "log-filename,log",
			Value: "../runtime/run/foo.log",
			Usage: "language for the greeting",
		},
	},
	Action: detach,
}

func detach(c *cli.Context) (err error) {
	_, err = os.Stat(c.String("p"))
	if err != nil {
		panic(fmt.Errorf("open pidfile failed, %s", err))
	}
	_, err = os.Stat(c.String("log"))
	if err != nil {
		panic(fmt.Errorf("open pidfile failed, %s", err))
	}
	cntxt := &daemon.Context{
		PidFileName: c.String("p"),
		LogFileName: c.String("log"),
		WorkDir:     "./",
		Args:        c.Args(),
	}
	ctx, cancel := context.WithCancel(context.Background())
	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()
	go Server(ctx)
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT) // 这里注册一下结束信号处理管道
	<-quit                                              // 无缓冲channel，在没有接收到信号之前后面的操作都会被阻塞你
	cancel()
	time.Sleep(time.Second * 1)
	return
}
