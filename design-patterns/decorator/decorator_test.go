package decorator

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

type RealHandler struct {
}

func (h *RealHandler) Handle(ctx context.Context) {
	fmt.Println("realHandler Handle")
	time.Sleep(1 * time.Second)
}

func TestDecorator(t *testing.T) {
	handler := LoggingDecorator{
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
		handler: &MetricsDecorator{
			handler: &RealHandler{},
		},
	}
	handler.Handle(context.Background())
}
