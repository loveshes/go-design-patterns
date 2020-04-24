# 6. 模板方法模式

**定义一个操作中的算法的骨架，将一些步骤延迟到子类中。模板方法使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。**

> 由于Go中没法想Java那样在父类中定义抽象方法，然后在子类中实现，以达到实际效果为在父类方法中去调用子类方法，这里我们使用传入func的方法来实现抽象方法。
>
> **如果有更好的实现方式，欢迎评论指出 \^_\^**

动物都有`Eat()`和`Sleep()`方法，针对不同的动物，Eat和Sleep方法的效果是不一样的。但是大致上是一样的，每种动物只是有细微差别，这时候我们就可以使用模板方法模式，抽象动物实现`模板方法`，在`模板方法`中给具体动物留下一个个的`空方法`，等具体动物去实现这里的`空方法`。具体操作如下

在路径`template\`下新建文件`template.go`，包名为`template`：

```go
package template

// ...
```

我们先定义动物接口和抽象动物，由于抽象动物不想被外界访问，需要首字母小写：

```go
// 动物接口
type Animal interface {
	Eat(food string) bool
	Sleep() bool
}

// 抽象动物
type animal struct {
	name string
	food []string
	max  int
	full bool
}
```

实现`eat()`方法的大致逻辑，这里的`eat()`是模板方法，不是给外界访问的，模板中留了`eat`、`full`、`refuse`三个空方法给具体的动物去实现：

```go
// 吃东西
// eat——还没吃饱的具体方法，full——吃了之后变饱的具体方法，refuse——拒绝吃东西的具体方法
// 返回是否吃了东西
// Go中只有组合，为了实现Animal中给具体的Pig/Dog留空方法，这里采用传入func的方式实现
func (a *animal) eat(food string, eat, full, refuse func() bool) bool {
	if !a.full {
		a.food = append(a.food, food)
		// 还没吃饱
		if len(a.food) < a.max {
			return eat() // 留给具体的动物去实现
		} else {
			a.full = true
			return full() // 留给具体的动物去实现
		}
	} else {
		return refuse() // 留给具体的动物去实现
	}
}
```

实现`sleep()`模板方法的大致逻辑，同样留了`full`、`notFull`两个空方法给具体的动物去实现：

```go
// 睡觉
// full——吃饱了的对应操作，notFull——没吃饱对应的操作
func (a *animal) sleep(full, notFull func() bool) bool {
	if a.full {
		return full() // 留给具体的动物去实现
	} else {
		return notFull() // 留给具体的动物去实现
	}
}
```

具体的动物`Dog`：

```go
// Dog
type Dog struct {
	*animal
}

func NewDog(name string, max int) *Dog {
	p := &animal{
		name: name,
		max:  max,
	}
	return &Dog{p}
}
```

`Dog`的`Eat()`方法，去实现具体的animal中eat方法的留空方法，这里定义了3个匿名函数去实现具体的逻辑：

```go
// Dog的Eat()方法，Dog吃饱了就不会继续吃
func (d *Dog) Eat(food string) bool {
	eat := func() bool { fmt.Printf("[%v]吃了[%v]，还要继续吃~\n", d.name, food); return true }
	full := func() bool { fmt.Printf("[%v]吃了[%v]，已经吃饱了~\n", d.name, food); return true }
	refuse := func() bool { fmt.Printf("[%v]吃饱了，拒绝吃东西~\n", d.name); return false }
	return d.animal.eat(food, eat, full, refuse)
}
```

`Dog`的`Sleep()`方法也类似：

```go
// Dog的Sleep()方法，Dog只有吃饱了才去睡觉
func (d *Dog) Sleep() bool {
	full := func() bool { fmt.Printf("[%v]吃饱了，去睡觉zzz\n", d.name); return true }
	notFull := func() bool { fmt.Printf("[%v]没吃饱，拒绝睡觉\n", d.name); return false }
	return d.animal.sleep(full, notFull)
}
```

具体的动物`Pig`：

```go
// Pig
type Pig struct {
	*animal
}

func NewPig(name string) *Pig {
	p := &animal{
		name: name,
	}
	return &Pig{p}
}
```

`Pig`的`Eat()`方法，实现animal中的留空方法，由于`Pig`不管有没有吃饱都要继续吃，所以只用实现一个`eat`就行

```go
// Pig的Eat()方法，Pig不管有没有吃饱都要继续吃
func (p *Pig) Eat(food string) bool {
	eat := func() bool { fmt.Printf("[%v]吃了[%v]，还要继续吃~\n", p.name, food); return true }
	return p.animal.eat(food, eat, eat, eat)
}
```

`Pig`的`Sleep()`方法也类似：

```go
// Pig的Sleep()方法，Pig不管有没有吃饱都可以去睡觉
func (p *Pig) Sleep() bool {
	full := func() bool { fmt.Printf("[%v]去睡觉zzz\n", p.name); return true }
	return p.animal.sleep(full, full)
}
```

在路径`template`的同级目录下新建`main.go`用于测试方法：

```go
package main

import (
	"fmt"
	"github.com/loveshes/go-design-patterns/pattern/template-method/template"
)

func main() {
	var animal template.Animal

	animal = template.NewPig("猪小妹")
	fmt.Println(animal.Eat("饲料"))
	fmt.Println(animal.Eat("饲料"))
	fmt.Println(animal.Sleep())
	fmt.Println()

	animal = template.NewDog("哈士奇", 2)
	fmt.Println(animal.Eat("骨头"))
	fmt.Println(animal.Sleep())
	fmt.Println(animal.Eat("肉"))
	fmt.Println(animal.Eat("大骨汤"))
	fmt.Println(animal.Sleep())
}
```

输出为

```
[猪小妹]吃了[饲料]，还要继续吃~
true
[猪小妹]吃了[饲料]，还要继续吃~
true
[猪小妹]去睡觉zzz
true

[哈士奇]吃了[骨头]，还要继续吃~
true
[哈士奇]没吃饱，拒绝睡觉
false
[哈士奇]吃了[肉]，已经吃饱了~
true
[哈士奇]吃饱了，拒绝吃东西~
false
[哈士奇]吃饱了，去睡觉zzz
true
```

可以看到，结果是符合预期了，Pig不管吃没吃饱都会去睡觉，Dog只有吃饱了才会去睡觉。

[完整示例](template/template.go)