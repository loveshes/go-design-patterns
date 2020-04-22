# 3. 代理模式

**代理模式，为其它对象提供一种代理以控制对这个对象的访问。**

比如小明向小红写情书，但是又不好意思直接给小红，于是就让小军替自己给小红。

在路径`proxy\`下新建文件`proxy.go`，包名为`proxy`：

```go
package proxy

// ...
```

定义小红和发送消息的接口：

```go
type Girl struct {
	Name string
}

// 发送消息的接口
type SendMsger interface {
	SendMsg(*Girl)
}
```

定义真实对象和真实对象的发送方法：

```go
// 真实对象
type TrueLove struct {
	Name string
}

// 真实对象的发送动作
func (p *TrueLove) SendMsg(g Girl) {
	fmt.Printf("[%s], 这是[%s]给你的情书\n", g.Name, p.Name)
}
```

定义代理对象、代理对象对真实对象的包装和代理对象的发送方法：

```go
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
```

在路径`proxy`的同级目录下新建`main.go`用于测试方法：

```go
package main

import (
	"fmt"
	"github.com/loveshes/go-design-patterns/pattern/proxy-pattern/proxy"
)

func main() {
	girl := proxy.Girl{"小红"}
	trueLove := proxy.TrueLove{"小明"}
	trueLove.SendMsg(girl)
	fmt.Println()

	proxy := proxy.Proxy{Name: "工具人小军"}
	proxy.Wrap(trueLove)
	proxy.SendMsg(girl)
}
```

输出为

```
[小红], 这是[小明]给你的情书

[小红], 这是[小明]给你的情书
——来自[工具人小军]
```

[完整示例](proxy/proxy.go)

