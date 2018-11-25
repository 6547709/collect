package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 0 {
		fmt.Println(os.Args[0]) // args 第一个参数是文件路径
		fmt.Println(filepath.Base(os.Args[0]))
	}
	//fmt.Println(os.Args[1]) // 第二个参数是，用户输入的参数，例如go run test1.go 123
}
