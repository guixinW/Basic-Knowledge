package adapter

import (
	"fmt"
	"testing"
)

var expect = "adaptee method"

type Client struct{}

func (c *Client) InsertLightningConnector(com Computer) {
	fmt.Println("Client plugs Lightning connector into computer.")
	com.InsertIntoLightingPort()
}

func TestAdapter(t *testing.T) {
	client := &Client{}
	windows := NewLG()
	adapter := NewUGREEN(windows)
	client.InsertLightningConnector(adapter)
}
