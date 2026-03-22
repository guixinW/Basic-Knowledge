package abstractfactory

import "testing"

// 安装并运行一套电脑系统外设，需要提供对应的生产线工厂
func SetupAndRunComputer(factory ComputerHardwareFactory) {
	mouse := factory.CreateMouse()
	keyboard := factory.CreateKeyboard()

	// 工厂通过统一的两条管线组装得到同品牌组件
	mouse.Click()
	keyboard.Type()
}

func TestComputerHardwareFactory(t *testing.T) {
	// 目前你只有罗技工厂，我们就指定要走罗技流水线
	var factory ComputerHardwareFactory = &LogitechFactory{}

	// 把这条专属于罗技的组装厂送进系统进行全套运转
	SetupAndRunComputer(factory)
}
