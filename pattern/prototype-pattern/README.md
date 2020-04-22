# 5. 原型模式

**原型模式，用原型实例指定创建对象的种类，并通过拷贝这些原型对象创建新的对象。**

比如有一份简历，其中个人信息部分应该是相同的，所以我们可以以此基础作为原型，然后通过原型的Clone()方法生成新的实例，在新的实例中修改特有的信息，以达到我们的目的。此外，不同的实例进行修改应该是**互不干扰**的。具体操作如下：

在路径`prototype\`下新建文件`prototype.go`，包名为`prototype`：

```go
package prototype

// ...
```

简历中预计包含姓名、性别、学校、申请公司、申请职位等信息，这里我们用一个结构体来保存学校信息：

```go
type School struct {
	Name    string
	Address string
	Level   string
}

// 创建School实例，返回对应指针
func NewSchool(name, address, level string) *School {
	return &School{
		Name:    name,
		Address: address,
		Level:   level,
	}
}
```

我们是通过原型的Clone()方法来拷贝一份新的实例的，所以为了不同实例之间能够互不干扰，我们应该使用深拷贝，School也要实现Clone()方法，这样原型对School()进行克隆的时候，就会调用School的Clone()方法。

```go
// School也要实现Clone()方法
func (s *School) Clone() *School {
	return NewSchool(s.Name, s.Address, s.Level)
}

// 重写School的String()方法，便于输出信息
func (s *School) String() string {
	return "{" + s.Name + " " + s.Address + " " + s.Level + "}"
}
```

然后就是Resume结构体和对应的新建Resume实例方法：

> 在这里需要注意的是，Go中结构体是值类型而不是指针类型，创建新的复合结构体时，会把结构体的值复制一份过去，所以这里使用School指针。此外也可以节省一定空间。

```go
type Resume struct {
	Name   string
	Gender string
	// 由于Go中结构体是值类型而不是指针类型，创建新的复合结构体时，会把值复制一份过去，所以这里使用School指针
	School        *School
	Apply4Company string
	Apply4Job     string
}

func NewResume(name, gender string, school *School, company, job string) *Resume {
	return &Resume{
		Name:          name,
		Gender:        gender,
		School:        school,
		Apply4Company: company,
		Apply4Job:     job,
	}
}
```

原型Clone()方法需要注意的点是，我们需要使用`r.School.Clone()`来把`r.School`（实际上是个指针）拷贝一份再放进去。

```go
// 原型的Clone()方法
func (r *Resume) Clone() *Resume {
	// 注意这里的r.School.Clone()
	return NewResume(r.Name, r.Gender, r.School.Clone(), r.Apply4Company, r.Apply4Job)
}
```

在路径`prototype`的同级目录下新建`main.go`用于测试方法：

```go
package main

import (
	"fmt"
	"github.com/loveshes/go-design-patterns/pattern/prototype-pattern/prototype"
)

func main() {
	ncu := prototype.NewSchool("南昌大学", "江西省南昌市", "211")
	proto := prototype.NewResume("王英俊", "男", ncu, "", "")

	// 简历一
	alibaba := proto.Clone()
	alibaba.Apply4Company = "Alibaba"
	alibaba.Apply4Job = "Java Web"
	fmt.Println("alibaba:", *alibaba)

	// 简历二
	bytedance := proto.Clone()
	// 修改复合结构体中的School.Level字段，看alibaba中的是否也会改变
	bytedance.School.Level = "双一流"
	bytedance.Apply4Company = "ByteDance"
	bytedance.Apply4Job = "Go"
	fmt.Println("修改School.Level后，alibaba:", *alibaba)
	fmt.Println("修改School.Level后，bytedance:", *bytedance)
}
```

输出为

```
alibaba: {王英俊 男 {南昌大学 江西省南昌市 211} Alibaba Java Web}
修改School.Level后，alibaba: {王英俊 男 {南昌大学 江西省南昌市 211} Alibaba Java Web}
修改School.Level后，bytedance: {王英俊 男 {南昌大学 江西省南昌市 双一流} ByteDance Go}
```

可以看到对简历二中School.Level的修改并没有影响到简历一。

此外可以把`r.Clone()`中的`r.School.Clone()`改成`r.School`，再看看有没有影响。

另外，由于Go中结构体本身就是值类型，把`Resume`结构体中的`*School`改成`School`，就算不使用`r.School.Clone()`，不同实例之间仍然互不干扰。

[完整示例](prototype/prototype.go)

