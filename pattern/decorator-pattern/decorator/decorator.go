package decorator

import (
	"fmt"
	"time"
)

// 工作代码
func Work() {
	fmt.Println("工作中...")
	time.Sleep(time.Second)
}

// logger函数包装工作代码
func Logger(f func()) {
	now := time.Now()
	fmt.Printf("开始:%v\n", now.Format("2006-01-02 15:04:05.000"))
	f() // 工作代码
	end := time.Now()
	fmt.Printf("结束:%v\n", end.Format("2006-01-02 15:04:05.000"))
	fmt.Printf("耗时:%v\n", end.Sub(now))
}

// 工作代码——需要接受参数
func WorkWithArgs(name string) {
	fmt.Printf("%s——工作中...", name)
	time.Sleep(time.Second * 2)
}

// 工作方法与装饰器不匹配，用中间方法进行桥接
func Bridge(f func(string), name string) func() {
	return func() {
		f(name)
	}
}
