package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func walDir(dir string, fileSize chan<- int64) {
	for _, entry := range scanDir(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			walDir(subDir, fileSize)
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

func main() {

	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{`e:\`}
	}

	fileSize := make(chan int64, 1)
	go func() {
		for _, dir := range roots {
			fmt.Println(dir)
			walDir(dir, fileSize)
		}
		close(fileSize)
	}()

	var files, bytes int64
	// go func() {
	// 	for size := range fileSize {
	// 		files++
	// 		bytes += size
	// 	}
	// }()

	for {
		select {
		case x := <-fileSize:
			files++
			bytes += x
		default:
			break
		}
	}

	fmt.Printf("%d files  %.1f GB\n", files, float64(bytes)/1e9)
}
