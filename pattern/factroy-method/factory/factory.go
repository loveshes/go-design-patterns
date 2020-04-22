package factory

// 【工厂方法模式】

import "fmt"

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

// 抽象工厂接口
type AnimalFactory interface {
	GetAnimal() Speaker
}

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
