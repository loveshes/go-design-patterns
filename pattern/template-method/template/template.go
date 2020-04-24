package template

// 【模板方法】

import "fmt"

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

// 睡觉
// full——吃饱了的对应操作，notFull——没吃饱对应的操作
func (a *animal) sleep(full, notFull func() bool) bool {
	if a.full {
		return full() // 留给具体的动物去实现
	} else {
		return notFull() // 留给具体的动物去实现
	}
}

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

// Dog的Eat()方法，Dog吃饱了就不会继续吃
func (d *Dog) Eat(food string) bool {
	eat := func() bool { fmt.Printf("[%v]吃了[%v]，还要继续吃~\n", d.name, food); return true }
	full := func() bool { fmt.Printf("[%v]吃了[%v]，已经吃饱了~\n", d.name, food); return true }
	refuse := func() bool { fmt.Printf("[%v]吃饱了，拒绝吃东西~\n", d.name); return false }
	return d.animal.eat(food, eat, full, refuse)
}

// Dog的Sleep()方法，Dog只有吃饱了才去睡觉
func (d *Dog) Sleep() bool {
	full := func() bool { fmt.Printf("[%v]吃饱了，去睡觉zzz\n", d.name); return true }
	notFull := func() bool { fmt.Printf("[%v]没吃饱，拒绝睡觉\n", d.name); return false }
	return d.animal.sleep(full, notFull)
}

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

// Pig的Eat()方法，Pig不管有没有吃饱都要继续吃
func (p *Pig) Eat(food string) bool {
	eat := func() bool { fmt.Printf("[%v]吃了[%v]，还要继续吃~\n", p.name, food); return true }
	return p.animal.eat(food, eat, eat, eat)
}

// Pig的Sleep()方法，Pig不管有没有吃饱都可以去睡觉
func (p *Pig) Sleep() bool {
	full := func() bool { fmt.Printf("[%v]去睡觉zzz\n", p.name); return true }
	return p.animal.sleep(full, full)
}
