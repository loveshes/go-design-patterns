package main

import "github.com/loveshes/go-design-patterns/pattern/factory-pattern/factory"

func main() {
	var af factory.AnimalFactory

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
