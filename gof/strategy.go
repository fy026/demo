package gof

type cashSuper interface {
	AcceptMoney(money float64) float64
}

//普通情况，没有折扣
type cashNormal struct {
}

func newCashNormal() cashNormal {
	instance := new(cashNormal)
	return *instance
}

func (c cashNormal) AcceptMoney(money float64) float64 {
	return money
}

//打折，传入打折的折扣，如0.8
type cashRebate struct {
	Rebate float64 //折扣
}

func newCashRebate(rebate float64) cashRebate {
	instance := new(cashRebate)
	instance.Rebate = rebate
	return *instance
}

func (c cashRebate) AcceptMoney(money float64) float64 {
	return money * c.Rebate
}

//直接返利，如满100返20
type cashReturn struct {
	MoneyCondition float64
	MoneyReturn    float64
}

func newCashReturn(moneyCondition float64, moneyReturn float64) cashReturn {
	instance := new(cashReturn)
	instance.MoneyCondition = moneyCondition
	instance.MoneyReturn = moneyReturn
	return *instance
}

func (c cashReturn) AcceptMoney(money float64) float64 {
	if money >= c.MoneyCondition {
		moneyMinus := int(money / c.MoneyCondition)
		return money - float64(moneyMinus)*c.MoneyReturn
	}
	return money
}

type CashContext struct {
	Strategy cashSuper
}

func NewCashContext(cashType string) CashContext {
	c := new(CashContext)
	//这里事实上是简易工厂模式的变形，用来生产策略
	switch cashType {
	case "打八折":
		c.Strategy = newCashRebate(0.8)
	case "满一百返20":
		c.Strategy = newCashReturn(100.0, 20.0)
	default:
		c.Strategy = newCashNormal()
	}
	return *c
}

//在策略生产成功后，我们就可以直接调用策略的函数。
func (c CashContext) GetMoney(money float64) float64 {
	return c.Strategy.AcceptMoney(money)
}
