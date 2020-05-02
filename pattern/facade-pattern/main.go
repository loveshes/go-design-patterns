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

	data := []interface{}{1, 2, 3, 4}
	list = facade.ArrayListOf(data...)
	list.RemoveLast()
	fmt.Println(list) // size=3, [1 2 3]
	fmt.Println(data) // [1 2 3 4]
}
