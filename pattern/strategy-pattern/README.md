# 1. 策略模式

**定义了策略家族，分别封装起来，让它们之间可以互相替换，此模式让策略的变化，不会影响到使用策略的客户（多态）。**实际上客户调用的是抽象的策略接口，不用关心策略的底层实现。

以超市打折为例，正常收费，打折收费和返利收费是具体策略，用一个策略接口去管理具体策略；同时还可以与简单工厂模式结合，传入策略名称，由工厂去创建策略管理类。

在路径`strategy\`下新建文件`strategy.go`，包名为`strategy`：

```go
package strategy

// ...
```

如下，首先定义了3个常量，用于规范具体的策略名称：

```go
// 3种具体的打折策略
const (
	ORIGINAL_PRICE    = iota // 0
	TEN_PERCENT_OFF          // 1
	FULL_100_MINUS_20        // 2
)
```

然后定义一个打折接口和实现打折接口的具体策略，打折接口Pricer和打折方法Price是要被外界调用的，所以需要首字母大写：

```go
// 打折接口
type Pricer interface {
	Price(cash float64) float64
}

// 原价
type normal struct {
}

// 打折
type discount struct {
	percent float64
}

// 每满fullMoney就减discountMoney
type reduction struct {
	fullMoney     float64
	discountMoney float64
}

// 均实现了pricer接口
func (n *normal) Price(cash float64) float64 {
	return cash
}
func (d *discount) Price(cash float64) float64 {
	return cash * d.percent
}
func (r *reduction) Price(cash float64) float64 {
	count := cash / r.fullMoney
	return cash - count*r.discountMoney
}
```

最后定义生成具体打折策略的工厂CashFactory，和工厂实现获取具体策略的方法GetCashPrice，由于均需要被外界调用，所以首字母需要大写：

```go
// 生成具体策略的工厂
type CashFactory struct {
}

// 将具体的策略实现封装起来，交由工厂来实现
func (c CashFactory) GetCashPrice(discountType int) Pricer {
	var cashPricer Pricer
	switch discountType {
	case ORIGINAL_PRICE:
		cashPricer = &normal{}
	case TEN_PERCENT_OFF:
		cashPricer = &discount{0.9}
	case FULL_100_MINUS_20:
		cashPricer = &reduction{100, 20}
	}
	return cashPricer
}

```

在路径`strategy`的同级目录下新建`main.go`用于测试方法：

```go
package main

import (
	"fmt"
	"github.com/loveshes/go-design-patterns/pattern/strategy-pattern/strategy"
)

func main() {
	var cf strategy.CashFactory
	var price strategy.Pricer
	price = cf.GetCashPrice(strategy.ORIGINAL_PRICE)
	fmt.Println(price.Price(210))

	price = cf.GetCashPrice(strategy.TEN_PERCENT_OFF)
	fmt.Println(price.Price(210))

	price = cf.GetCashPrice(strategy.FULL_100_MINUS_20)
	fmt.Println(price.Price(210))
}
```

输出为

```
210
189
168
```

[完整示例](strategy/strategy.go)

