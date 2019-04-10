package cash

import "errors"

type CashI interface {
	AcceptCash(money float64) float64
}

type CashNormal struct {
}

func (cn *CashNormal) AcceptCash(money float64) float64 {
	return money
}

type CashRebate struct {
	moneyRebate float64
}

func (cr *CashRebate) AcceptCash(money float64) float64 {
	return money * cr.moneyRebate
}

type CashReturn struct {
	moneyCondition float64
	moneyReturn    float64
}

func (cr *CashReturn) AcceptCash(money float64) float64 {
	result := money
	if money >= cr.moneyCondition {
		result = money - (money/cr.moneyCondition)*cr.moneyReturn
	}
	return result
}

type CashContext struct {
	cs CashI
}

func (cc *CashContext) GetResult(money float64) float64 {
	return cc.cs.AcceptCash(money)
}

func NewCashContext(ctype string) (*CashContext, error) {
	switch ctype {
	case "正常收费":
		return &CashContext{&CashNormal{}}, nil
	case "满300返100":
		return &CashContext{&CashReturn{300, 100}}, nil
	case "打8折":
		return &CashContext{&CashRebate{0.8}}, nil
	default:
		return nil, errors.New("not has this type")
	}
}
