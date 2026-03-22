package strategy

type TaxBase int

const (
	CNTax TaxBase = iota
	USTax
	DETax
	FRTax
)

type SalesOrderOld struct {
	tax TaxBase
}

// CalculateTax 每次添加不同的Tax计算方法都需要修改CalculateTax本体，不符合设计模式中的“修改封闭，扩展开放”的原则。
func (s *SalesOrderOld) CalculateTax() float64 {
	var tax float64
	if s.tax == CNTax {

	} else if s.tax == USTax {

	} else if s.tax == DETax {

	} else if s.tax == FRTax {

	} else {

	}
	return tax
}
