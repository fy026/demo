package gof

import (
	"log"
	"testing"
)

func TestStrategy(t *testing.T) {
	money := 100.0
	cc := NewCashContext("打八折")
	money = cc.GetMoney(money)
	log.Println("100打八折实际金额为", money)

	money = 199
	cc = NewCashContext("满一百返20")
	money = cc.GetMoney(money)
	log.Println("199满一百返20实际金额为", money)

	money = 199
	cc = NewCashContext("没有折扣")
	money = cc.GetMoney(money)
	log.Println("199没有折扣实际金额为", money)
}
