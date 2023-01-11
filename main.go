package main

import "github.com/agentwyz/testmod/lib"

func main() {
	//直接使用go mod tidy进行导入
	lib.Hello("wyz")
}