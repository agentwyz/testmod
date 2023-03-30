package main

import (
	"github.com/agentwyz/testmod/lib"
	"fmt"
)



func main() {
	//直接使用go mod tidy进行导入
	lib.Hello("wyz")
	test()
}

//测试分支
func test() {
	fmt.Println("test this brach")
}