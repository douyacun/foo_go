package main

import (
	"fmt"
	"gopl.io/ch8/thumbnail"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	dir := "/Users/liuning/Documents/github/foo/asserts/images/"
	outDir := "/Users/liuning/Documents/github/foo/asserts/thumbs/"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("ioutil read file failed, %s", err)
		return
	}
	if err := os.RemoveAll(outDir); err != nil {
		fmt.Printf("remove out dir faield, %s", err)
		return
	}
	if err = os.Mkdir(outDir, 0777); err != nil {
		fmt.Printf("mkdir failed, %s", err)
	}
	makeThumbnails(files, dir, outDir)
}

func makeThumbnails(files []os.FileInfo, dir string, outDir string) {
	wg := sync.WaitGroup{}
	type item struct {
		thumbfile string
		err       error
	}
	sizes := make(chan int64)
	var (
		ext string
	)
	for _, f := range files {
		wg.Add(1)
		go func(f os.FileInfo) {
			defer wg.Done()
			file := dir + f.Name()
			ext = filepath.Ext(f.Name())
			switch ext {
			case ".jpg", ".jpeg", ".png":
				goto Make
			default:
				return
			}
		Make:
			thumbfile := outDir + strings.TrimSuffix(f.Name(), ext) + ".thumb" + ext
			err := thumbnail.ImageFile2(thumbfile, file)
			if err != nil {
				fmt.Printf("make thumb image failed, %s\n", err)
				return
			}
			fileInfo, err := os.Stat(thumbfile)
			if err != nil {
				fmt.Printf("%s\n", err)
				return
			}
			sizes <- fileInfo.Size()
		}(f)
	}
	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for s := range sizes {
		total += s
	}
	fmt.Printf("%dkB\n", total/(1024))
}
