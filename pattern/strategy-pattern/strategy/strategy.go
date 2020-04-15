package strategy

//【策略模式】

// 3种具体的打折策略
const (
	ORIGINAL_PRICE    = iota // 0
	TEN_PERCENT_OFF          // 1
	FULL_100_MINUS_20        // 2
)

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
