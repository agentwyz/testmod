package main

import (
	"fmt"
	"github.com/agentwyz/testmod/lib"
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

//看看是否冲突
func ddd() {
	fmt.Println("test dddd")
}


//测试分支冲突
func test01() {
	fmt.Println("test this brach01")
}

//再次修改
func test02() {
	fmt.Println("test this branch02")
}
