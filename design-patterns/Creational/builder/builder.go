package builder

import (
	"errors"
)

/*-----------传统Builder----------------*/
type ComputerConfiguration struct {
	CPU         string
	GPU         string
	Motherboard string
	Memory      int
	Disk        int
}

type ComputerBuilder struct {
	config *ComputerConfiguration
	err    error
}

func (b *ComputerBuilder) Build() (*ComputerConfiguration, error) {
	if b.err != nil {
		return nil, b.err
	}
	return b.config, nil
}

func NewComputerBuilder(cpu string) *ComputerBuilder {
	return &ComputerBuilder{
		config: &ComputerConfiguration{
			CPU:         cpu,
			GPU:         "RTX 4060",
			Memory:      16,
			Disk:        1024,
			Motherboard: "B650",
		},
	}
}

func (b *ComputerBuilder) WithCPU(cpu string) *ComputerBuilder {
	b.config.CPU = cpu
	return b
}

func (b *ComputerBuilder) WithGPU(gpu string) *ComputerBuilder {
	if gpu != "AMD" && gpu != "NVIDIA" {
		b.err = errors.New("invalid GPU")
		return b
	}
	b.config.GPU = gpu
	return b
}

func (b *ComputerBuilder) WithMotherBoard(motherboard string) *ComputerBuilder {
	b.config.Motherboard = motherboard
	return b
}

func (b *ComputerBuilder) WithMemory(memory int) *ComputerBuilder {
	b.config.Memory = memory
	return b
}

func (b *ComputerBuilder) WithDisk(disk int) *ComputerBuilder {
	b.config.Disk = disk
	return b
}

/*-----------函数式选项模式----------------*/
type ComputerConfiguration2 struct {
	CPU         string
	GPU         string
	Motherboard string
	Memory      int
	Disk        int
}

type ComputerOption func(*ComputerConfiguration2)

func WithCPU(cpu string) ComputerOption {
	return func(c *ComputerConfiguration2) {
		c.CPU = cpu
	}
}

func WithGPU(gpu string) ComputerOption {
	return func(c *ComputerConfiguration2) {
		c.GPU = gpu
	}
}

func WithMotherBoard(motherboard string) ComputerOption {
	return func(c *ComputerConfiguration2) {
		c.Motherboard = motherboard
	}
}

func WithMemory(memory int) ComputerOption {
	return func(c *ComputerConfiguration2) {
		c.Memory = memory
	}
}

func WithDisk(disk int) ComputerOption {
	return func(c *ComputerConfiguration2) {
		c.Disk = disk
	}
}

func NewComputer(cpu string, opts ...ComputerOption) *ComputerConfiguration2 {
	config := &ComputerConfiguration2{
		CPU:         cpu,
		GPU:         "AMD",
		Motherboard: "MSI",
		Memory:      24,
		Disk:        512,
	}
	for _, opt := range opts {
		opt(config)
	}
	return config
}
