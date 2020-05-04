package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)
var sema = make(chan struct{}, 20)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)

	var n sync.WaitGroup
	go func() {
		n.Add(1)
		for _, entry := range roots {
			go walkDir(entry, fileSizes, &n)
		}
	}()
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	var nfiles, nbytes int64
	var tick <-chan time.Time
	tick = time.Tick(500 * time.Millisecond)
loop:
	for {
		select {
		case <-tick:
			fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/(1<<30))
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		}
	}

	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/(1<<30))
}

func walkDir(dir string, fileSizes chan<- int64, n *sync.WaitGroup) {
	defer n.Done()
	entries := dirEntries(dir)
	for _, entry := range entries {
		if entry.IsDir() {
			n.Add(1)
			go walkDir(filepath.Join(dir, entry.Name()), fileSizes, n)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirEntries(d string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	entries, err := ioutil.ReadDir(d)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return entries
}
