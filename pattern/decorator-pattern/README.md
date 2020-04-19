# 2. 装饰器模式

**装饰器模式，动态地给一个对象添加一些额外的职责，就增加功能来说，装饰器模式比生成子类更加灵活。它把每个要装饰的功能放在单独的类中，并让这个类包装它所要装饰的对象，在使用时要注意装饰的顺序。**

比如我们想给核心代码添加日志打印功能，但是又不能改动原有代码，可以使用装饰器模式来包装原有的代码。

在路径`decorator\`下新建文件`decorator.go`，包名为`decorator`：

```go
package decorator

// ...
```

如下为工作代码：

```go
// 工作代码
func Work() {
	fmt.Println("工作中...")
	time.Sleep(time.Second)
}
```

定义一个装饰器函数来包装工作代码：

```go
// logger函数包装工作代码
func Logger(f func()) {
	now := time.Now()
	fmt.Printf("开始:%v\n", now.Format("2006-01-02 15:04:05.000"))
	f() // 工作代码
	end := time.Now()
	fmt.Printf("结束:%v\n", end.Format("2006-01-02 15:04:05.000"))
	fmt.Printf("耗时:%v\n", end.Sub(now))
}
```

加入工作代码的函数签名与Logger()函数的参数不匹配，这时候可以使用桥接模式定义一个桥接函数，来把2个函数桥接起来：

```go
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
```

在路径`decorator`的同级目录下新建`main.go`用于测试方法：

```go
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
```

[完整示例](decorator/decorator.go)

