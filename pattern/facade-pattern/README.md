# 外观模式

**外观模式，为子系统中的一组接口提供一个一致的界面，此模式定义了一个高层接口，这个接口使得这一子系统更加容易使用。**

像平时我们操作slice的时候，尤其是往slice中间插入、删除元素多有不便，我们可以在slice的外面套一层，使其更易使用。此外，由于使用了外观模式，使用者无需关心其内部实现，可以实现ArrayList和LinkedList的随心切换。

在路径`facade\`下新建文件`facade.go`，包名为`facade`：

```go
package facade

// ...
```

我们首先有一个统一的List接口，接下来的不管是ArrayList还是LinkedList都会实现该接口：

```go
// 统一接口
type List interface {
	Add(idx int, e interface{}) bool
	AddLast(e interface{}) bool
	Remove(idx int) (interface{}, bool)
	RemoveLast() (interface{}, bool)
	Get(idx int) (interface{}, bool)
}
```

然后有ArrayList的结构和构造方法，需要注意capacity的处理：

```go
type ArrayList struct {
	ele  []interface{}
	size int
}

func NewArrayList(capacity int) *ArrayList {
	if capacity < 16 {
		capacity = 16
	}
	return &ArrayList{
		ele:  make([]interface{}, 0, capacity),
		size: 0,
	}
}
```

此外，我们也可以根据传入的元素来创建一个ArrayList，这是构建外观的核心：

```go
func ArrayListOf(ele ...interface{}) *ArrayList {
	return &ArrayList{
		ele:  ele,
		size: len(ele),
	}
}
```

接下来就是实现List接口的各种方法了，尤其要注意往中间插入元素和删除元素的处理：

```go
func (a *ArrayList) Add(idx int, e interface{}) bool {
	if idx > a.size {
		return false
	}
	a.ele = append(a.ele, nil)
	copy(a.ele[idx+1:], a.ele[idx:]) // 后移一位, 把idx位置空出来
	a.ele[idx] = e
	a.size++
	return true
}

func (a *ArrayList) AddLast(e interface{}) bool {
	a.ele = append(a.ele, e)
	a.size++
	return true
}

func (a *ArrayList) Remove(idx int) (interface{}, bool) {
	if idx >= a.size {
		return nil, false
	}
	e := a.ele[idx]
	a.ele = append(a.ele[:idx], a.ele[idx+1:]...) // 跳过idx元素
	a.size--
	return e, true
}

func (a *ArrayList) RemoveLast() (interface{}, bool) {
	e := a.ele[a.size-1]
	a.ele = append(a.ele[:a.size-1])
	a.size--
	return e, true
}

func (a *ArrayList) Get(idx int) (interface{}, bool) {
	if idx >= a.size {
		return nil, false
	}
	e := a.ele[idx]
	return e, true
}

func (a *ArrayList) String() string {
	return fmt.Sprintf("size=%v, %v", a.size, a.ele)
}
```

在路径`facade`的同级目录下新建`main.go`用于测试方法：

```go
package main

import (
	"fmt"
	"github.com/loveshes/go-design-patterns/pattern/facade-pattern/facade"
)

func main() {
	var list facade.List
	list = facade.NewArrayList(10)
	list.AddLast(1)
	list.AddLast(3)
	list.AddLast(4)
	list.Add(1, 2)
	fmt.Println(list)              // size=4, [1 2 3 4]
	fmt.Println(list.RemoveLast()) // 4 true
	fmt.Println(list)              // size=3, [1 2 3]
	fmt.Println(list.Remove(0))    // 1 true
	fmt.Println(list)              // size=2, [2 3]
	fmt.Println(list.Get(1))       // 3 true

	list = facade.ArrayListOf(1, 2, 3, 4, 5, 6)
	fmt.Println(list)              // size=6, [1 2 3 4 5 6]
	fmt.Println(list.AddLast(7))   // true
	fmt.Println(list)              // size=7, [1 2 3 4 5 6 7]
	fmt.Println(list.Add(1, 100))  // true
	fmt.Println(list)              // size=8, [1 100 2 3 4 5 6 7]
	fmt.Println(list.Remove(1))    // 100 true
	fmt.Println(list)              // size=7, [1 2 3 4 5 6 7]
	fmt.Println(list.RemoveLast()) // 7 true
	fmt.Println(list)              // size=6, [1 2 3 4 5 6]
	fmt.Println(list.Get(2))       // 3 true
}
```

输出如上，结果完全符合预期。

需要注意的是，使用如下方式构建ArrayList，当我们对ArrayList进行操作时，传入的data本身并不会改变，这并不代表我们在构建ArrayList的时候拷贝了data内部的值，而且我们在ArrayList内部进行操作时，经常使用到的`append()`函数会重新指向`slice`，当发生扩容时，会构建一份新的底层数组，故不一定会影响到原`slice`：

```go
data := []interface{}{1, 2, 3, 4}
list = facade.ArrayListOf(data...)
list.RemoveLast()
fmt.Println(list) // size=3, [1 2 3]
fmt.Println(data) // [1 2 3 4]
```

[完整示例](facade/facade.go)

