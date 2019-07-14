package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 递归遍历输入的目录的所有文件总数和其大小
func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSize := make(chan int64)
	go func() {
		for _, root := range roots {
			travel(root, fileSize)
		}
		// count size over, close the channel
		close(fileSize)
	}()


	// show the result
	var nfiles, nbytes int64

	// receive file size data from fileSize channel before it is closed
	for size := range fileSize {
		nfiles++
		nbytes += size
	}

	// print result
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func travel(dir string, fileSize chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			travel(subdir, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read dir err: %v\n", err)
		return nil
	}
	return entries
}
