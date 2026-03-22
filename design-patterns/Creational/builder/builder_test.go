package builder

import (
	"fmt"
	"testing"
)

func TestTraditionalBuilder(t *testing.T) {
	computer, err := NewComputerBuilder("AMD").WithMotherBoard("ASUS").WithGPU("NVIDIA").WithMemory(16).WithDisk(1024).Build()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(computer)
}

func TestFunctionalOptions(t *testing.T) {
	computer := NewComputer("Intel", WithDisk(2048))
	fmt.Println(computer)
}
