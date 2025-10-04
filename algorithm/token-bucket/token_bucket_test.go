package token_bucket

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewTokenBucket(t *testing.T) {
	tokenBucket := NewTokenBucket(5, 3)
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			tokenBucket.Take()
			fmt.Printf("%v get token\n", id)
		}(i)
	}
	wg.Wait()
	end := time.Now().Sub(start)
	fmt.Printf("task finished, took time %v \n", end)
}
