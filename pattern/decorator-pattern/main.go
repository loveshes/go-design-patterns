package main

import (
	"fmt"
	"github.com/loveshes/go-design-patterns/pattern/decorator-pattern/decorator"
)

func main() {
	work := decorator.Work
	// 直接调用work()
	work()
	fmt.Println()

	// 使用装饰器调用work()
	decorator.Logger(work)
	fmt.Println()

	// 桥接方法
	work2 := decorator.WorkWithArgs
	bridge := decorator.Bridge
	decorator.Logger(bridge(work2, "[工作方法]"))
}
