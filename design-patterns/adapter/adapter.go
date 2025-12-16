package adapter

import "fmt"

/*Computer Target: 目标接口*/
type Computer interface {
	InsertIntoLightingPort()
}

type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Mac Need Light Port")
}

/*Window Adaptee: 被适配接口*/
type Window interface {
	InsertIntoHDMIPort()
}

// LG 屏幕实现

func NewLG() *LG {
	return &LG{}
}

type LG struct {
}

func (l *LG) InsertIntoHDMIPort() {
	fmt.Println("LG Need HDMI Port")
}

/*Adapter 转换器*/
type Adapter interface {
}

// UGREEN 转换器实现
func NewUGREEN(mechine Window) *UGREEN {
	return &UGREEN{
		windowMachine: mechine,
	}
}

type UGREEN struct {
	windowMachine Window
}

func (u *UGREEN) InsertIntoLightingPort() {
	fmt.Println("UGREEN transform HDMI signal to light port signla")
	u.windowMachine.InsertIntoHDMIPort()
}
