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
