package strategy

// TaxStrategy 定义税率接口
type TaxStrategy interface {
	Calculate(order *SalesOrderNew) float64
}

type CNTaxImplementation struct{}

func (*CNTaxImplementation) Calculate(order *SalesOrderNew) float64 {
	var tax float64
	return tax
}

type USTaxImplementation struct{}

func (*USTaxImplementation) Calculate(order *SalesOrderNew) float64 {
	var tax float64
	return tax
}

type DETaxImplementation struct{}

func (*DETaxImplementation) Calculate(order *SalesOrderNew) float64 {
	var tax float64
	return tax
}

type FRTaxImplementation struct{}

func (*FRTaxImplementation) Calculate(order *SalesOrderNew) float64 {
	var tax float64
	return tax
}

// SalesOrderNew 订单信息
type SalesOrderNew struct {
	strategy TaxStrategy
}

func NewSalesOrderNew(strategy TaxStrategy) *SalesOrderNew {
	return &SalesOrderNew{
		strategy: strategy,
	}
}

// CalculateTax 根据工厂函数传入的Strategy选择不同的税率计算方法
func (s *SalesOrderNew) CalculateTax() float64 {
	return s.strategy.Calculate(s)
}
