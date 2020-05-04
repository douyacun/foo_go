package commands

import (
	"fmt"
	"github.com/urfave/cli"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var CurlCommand = cli.Command{
	Name:   "curl",
	Usage:  "send a http request",
	Action: CurlAction,
}

func CurlAction(ctx *cli.Context) error {
	ch := make(chan string)
	for _, url := range ctx.Args() {
		go fetch(url, ch)
	}
	for range ctx.Args() {
		fmt.Println(<-ch)
	}
	return nil
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("http get %s failed, %s", url, err)
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("request[%s] response.body read failed, %s", url, err)
		return
	}
	_ = resp.Body.Close()

	latency := time.Now().Sub(start)
	ch <- fmt.Sprintf("%s %7d %s", latency, nbytes, url)
}
