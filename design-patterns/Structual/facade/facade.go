package facade

import "fmt"

func NewAPI() API {
	return ApiImpl{
		serviceA: NewServiceA(),
		serviceB: NewServiceB(),
	}
}

type API interface {
	Call() string
}

type ApiImpl struct {
	serviceA ServiceA
	serviceB ServiceB
}

func (impl ApiImpl) Call() string {
	aRet := impl.serviceA.CallA()
	bRet := impl.serviceB.CallB()
	return fmt.Sprintf("%s, %s", aRet, bRet)
}

type ServiceA interface {
	CallA() string
}

type serviceAImpl struct {
}

func NewServiceA() ServiceA {
	return &serviceAImpl{}
}

func (service *serviceAImpl) CallA() string {
	return "A"
}

type serviceBImpl struct {
}

func NewServiceB() ServiceB {
	return &serviceBImpl{}
}

type ServiceB interface {
	CallB() string
}

func (service *serviceBImpl) CallB() string {
	return "B"
}
