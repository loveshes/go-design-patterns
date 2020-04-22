package main

import "github.com/loveshes/go-design-patterns/pattern/factroy-method/factory"

func main() {
	// 抽象工厂
	var af factory.AnimalFactory

	// 狗工厂
	af = &factory.DogFactory{} // 由于实现AnimalFactory接口的接收者都是指针类型，所以是对应的指针实现了接口，这里需要加上&
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
