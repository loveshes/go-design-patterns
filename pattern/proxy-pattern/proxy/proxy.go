package proxy

// 【代理模式】

import "fmt"

type Girl struct {
	Name string
}

// 发送消息的接口
type SendMsger interface {
	SendMsg(*Girl)
}

// 真实对象
type TrueLove struct {
	Name string
}

// 真实对象的发送动作
func (p *TrueLove) SendMsg(g Girl) {
	fmt.Printf("[%s], 这是[%s]给你的情书\n", g.Name, p.Name)
}

// 代理对象
type Proxy struct {
	Name string
	TrueLove
}

// 代理对象对真实对象进行包装
func (p *Proxy) Wrap(t TrueLove) {
	p.TrueLove = t
}

// 代理对象的发送动作
func (p *Proxy) SendMsg(g Girl) {
	p.TrueLove.SendMsg(g) // 调用真实对象的方法
	fmt.Printf("——来自[%v]\n", p.Name)
}
