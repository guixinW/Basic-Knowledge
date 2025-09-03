package decorator

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Handler interface {
	Handle(ctx context.Context)
}

type LoggingDecorator struct {
	handler Handler
	logger  *log.Logger
}

func (d *LoggingDecorator) Handle(ctx context.Context) {
	d.logger.Println("开始执行")
	d.handler.Handle(ctx)
}

type MetricClient struct {
}

func (m *MetricClient) Record(key string, value interface{}) {
	fmt.Printf("%v : %v", key, value)
}

type MetricsDecorator struct {
	handler      Handler
	metricClient MetricClient
}

func (m *MetricsDecorator) Handle(ctx context.Context) {
	start := time.Now()
	defer func() {
		end := time.Since(start)
		m.metricClient.Record("执行时间", end)
	}()
	m.handler.Handle(ctx)
}
