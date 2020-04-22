# 4. 工厂方法模式

**工厂模式方法，定义一个用于创建对象对象的接口，让子类决定实例化哪一个类，工厂方法使一个类的实例化延迟到其子类。**

之前我们使用简单工厂模式是一个工厂负责创建各种动物，如果需要新增动物的话，工厂需要修改原有代码逻辑。这里使用工厂方法模式，不再由一个工厂负责创建所有动物，而且由一堆实现了工厂接口的具体动物工厂去创建具体的动物，每个工厂只能创建自己的动物。具体操作如下：

在路径`factory\`下新建文件`factory.go`，包名为`factory`：

```go
package factory

// ...
```

前面的代码与简单工厂模式相同：

```go
// 动物
type animal struct {
	name   string
	master string
}

// 给动物设置基本信息
func (a *animal) SetInfo(name, master string) {
	a.name = name
	a.master = master
}

// 具体物种
type dog struct {
	animal
}
type cat struct {
	animal
}
type pig struct {
	animal
}

// 动物都有Speak()方法
type Speaker interface {
	Speak()
	SetInfo(name, master string) // 这里调用的实际上是animal类型的方法
}

func (d *dog) Speak() {
	fmt.Printf("我叫[%v], 我的主人是[%v], 汪汪汪~\n", d.name, d.master)
}

func (c *cat) Speak() {
	fmt.Printf("我叫[%v], [%v]才不是我的主人呢, 明明是我的铲屎官\n", c.name, c.master)
}

func (p *pig) Speak() {
	fmt.Printf("我叫[%v], 每天都是[%v]给饭我吃, 我吃饱了就想睡觉\n", p.name, p.master)
}
```

这里要注意了，有一个抽象工厂接口：

```go
// 抽象工厂接口
type AnimalFactory interface {
	GetAnimal() Speaker
}
```

这个抽象工厂接口有不同的实现，每个实现都是一个创建具体动物的工厂：

```go
// 生产具体动物的具体工厂
type DogFactory struct{}
type CatFactory struct{}
type PigFactory struct{}

// 狗工厂
func (d *DogFactory) GetAnimal() Speaker {
	return &dog{}
}

// 猫工厂
func (c *CatFactory) GetAnimal() Speaker {
	return &cat{}
}

// 猪工厂
func (p *PigFactory) GetAnimal() Speaker {
	return &pig{}
}
```

在路径`factory`的同级目录下新建`main.go`用于测试方法：

```go
package main

import "github.com/loveshes/go-design-patterns/pattern/factroy-method/factory"

func main() {
	// 抽象工厂
	var af factory.AnimalFactory

	// 狗工厂
    // 由于实现AnimalFactory接口的接收者都是指针类型，所以是对应的指针实现了接口，这里需要加上&
	af = &factory.DogFactory{} 
	dog := af.GetAnimal()
	dog.SetInfo("哈士奇", "张三")
	dog.Speak()

	// 猫工厂
	af = &factory.CatFactory{}
	cat := af.GetAnimal()
	cat.SetInfo("英短", "李四")
	cat.Speak()

	// 猪工厂
	af = &factory.PigFactory{}
	pig := af.GetAnimal()
	pig.SetInfo("猪英俊", "王五")
	pig.Speak()
}
```

输出为

```
我叫[哈士奇], 我的主人是[张三], 汪汪汪~
我叫[英短], [李四]才不是我的主人呢, 明明是我的铲屎官
我叫[猪英俊], 每天都是[王五]给饭我吃, 我吃饱了就想睡觉
```

[完整示例](factory/factory.go)

