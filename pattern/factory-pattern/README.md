# 0. 简单工厂模式

**通过传入参数给工厂，让工厂来生成我们想要的实际对象，进而对对象进行一系列操作。**

这里想要实现传入动物的类型名称，让工厂来生成我们想要的具体动物。具体操作如下：

这里我们定义了动物这种类型`animal`和给动物设置基本信息的方法`SetInfo(name, master string)`，注意SetInfo()方法是要向外暴露的，所有需要首字母大写。

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
```

然后定义具体的物种以及实际供外部调用的`Speak()`方法，要注意的是，Speaker接口中的SetInfo()方法是animal类型中的，由于各个物种都是包含了animal，所以自然也包含了animal类型的方法：

```go
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

func (d dog) Speak() {
	fmt.Printf("我叫[%v], 我的主人是[%v], 汪汪汪~\n", d.name, d.master)
}

func (c cat) Speak() {
	fmt.Printf("我叫[%v], [%v]才不是我的主人呢, 明明是我的铲屎官\n", c.name, c.master)
}

func (p pig) Speak() {
	fmt.Printf("我叫[%v], 每天都是[%v]给饭我吃, 我吃饱了就想睡觉\n", p.name, p.master)
}
```

最后定义了生产动物的工厂`AnimalFactory`和获取动物的方法`GetAnimal(animalType string) Speaker`，工厂和获取动物的方法都是需要向外暴露的，首字母大写。

```go
// 生产动物的工厂
type AnimalFactory struct{}

// 工厂生产动物
func (a AnimalFactory) GetAnimal(animalType string) Speaker {
	var speaker Speaker
	switch animalType {
	case "dog":
		speaker = &dog{}
	case "cat":
		speaker = &cat{}
	case "pig":
		speaker = &pig{}
	}
	return speaker
}
```

然后我们就可以使用了：

```go
func main() {
	var af AnimalFactory

	dog := af.GetAnimal("dog")
	dog.SetInfo("小黑", "张三")
	dog.Speak()

	cat := af.GetAnimal("cat")
	cat.SetInfo("喵呜", "marry")
	cat.Speak()

	pig := af.GetAnimal("pig")
	pig.SetInfo("粉红猪小妹", "猪倌")
	pig.Speak()
}
```

[完整示例](factory/factory.go)