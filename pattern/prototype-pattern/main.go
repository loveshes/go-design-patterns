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
