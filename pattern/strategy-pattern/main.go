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
