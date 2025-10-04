package rwlock

import (
	"fmt"
	"sync"
	"testing"
)

func write() {
	fmt.Println("write")
}

func read() {
	fmt.Println("read")
}

func TestSyncRWMutex(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			read()
		}()
	}
	wg.Wait()
}
