package token_bucket

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity      float64
	rate          float64
	tokens        float64
	lastTokenTime time.Time
	mu            sync.Mutex
}

func NewTokenBucket(capacity, rate float64) *TokenBucket {
	return &TokenBucket{capacity: capacity, rate: rate, lastTokenTime: time.Now(), tokens: capacity}
}

func (tb *TokenBucket) refill() {
	now := time.Now()
	duration := now.Sub(tb.lastTokenTime)
	tokensToAdd := duration.Seconds() * tb.rate

	tb.tokens += tokensToAdd
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}

	tb.lastTokenTime = now
}

func (tb *TokenBucket) Take() {
	for {
		tb.mu.Lock()
		tb.refill()
		if tb.tokens > 1 {
			tb.tokens -= 1
			tb.mu.Unlock()
			return
		}
		requireTokens := 1 - tb.tokens
		durationWait := time.Duration(requireTokens / tb.rate * float64(time.Second))
		tb.mu.Unlock()
		time.Sleep(durationWait)
	}
}
