package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walDir(dir string, fileSize chan<- int64, wg *sync.WaitGroup) {
	for _, entry := range scanDir(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go walDir(subDir, fileSize, wg)
		} else {
			info, err := entry.Info()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fileSize <- info.Size()
		}
	}
}

func scanDir(dir string) []fs.DirEntry {
	fs, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
	}

	return fs
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {

	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{`e:\phpenv`}
	}

	var wg sync.WaitGroup
	fileSize := make(chan int64, 1)
	for _, dir := range roots {
		wg.Add(1)
		go walDir(dir, fileSize, &wg)
	}

	go func() {
		wg.Wait()
		close(fileSize)
	}()

	var files, bytes int64
	go func() {
		for size := range fileSize {
			files++
			bytes += size
		}
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
loop:
	for {
		select {
		case x, ok := <-fileSize:
			if !ok {
				break loop
			}
			files++
			bytes += x
		case <-tick:
			fmt.Printf("%d files  %.1f GB\n", files, float64(bytes)/1e9)
		}
	}

}
