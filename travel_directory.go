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
	var fileNumber, bytes int64

	// receive file size data from fileSize channel before the channel is closed
	for size := range fileSize {
		fileNumber++
		bytes += size
	}

	// print result
	fmt.Printf("%d files  %.1f GB\n", fileNumber, float64(bytes)/1e9)
}

// 递归遍历传入的路径
func travel(dir string, fileSize chan<- int64) {
	for _, entry := range directs(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			travel(subDir, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

// 根据传入的路径获取目录信息
func directs(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read dir err: %v\n", err)
		return nil
	}
	return entries
}
