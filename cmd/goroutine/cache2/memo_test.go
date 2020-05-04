package cache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMemo_Get(t *testing.T) {
	cache := New(HttpGetBody)
	incomingURLs := []string{
		"https://www.00h.tv",
		"https://www.baidu.com",
		"https://www.douban.com",
		"https://www.00h.tv",
		"https://www.douban.com",
		"https://www.baidu.com",
	}
	wg := sync.WaitGroup{}
	for _, url := range incomingURLs {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			defer wg.Done()
			value, err := cache.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%s\t%s\t%d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	wg.Wait()
}
