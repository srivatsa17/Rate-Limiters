package rate_limiters

import (
	"fmt"
	"sync"
	"time"
)

// How does Token Bucket Work ?
// 1.	Consider a bucket having many tokens
// 2.	The bucket will have a capacity beyond which tokens can't be held.
// 3.	The tokens are refilled at a certain rate into the bucket i.e. for eg 10 tokens / second.
// 4.	If there are enough tokens in the bucket, the request is accepted, else denied.

type TokenBucket struct {
	capacity       int
	rate           int
	tokens         int
	mu             sync.Mutex
	lastRefilledAt time.Time
}

func NewTokenBucket(capacity, rate int) *TokenBucket {
	tb := &TokenBucket{
		capacity:       capacity,
		tokens:         capacity,
		rate:           rate,
		lastRefilledAt: time.Now(),
	}

	return tb
}

func (tb *TokenBucket) AllowRequest() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	elapsed := time.Since(tb.lastRefilledAt).Seconds()
	refillTokens := int(elapsed) * tb.rate
	if refillTokens > 0 {
		tb.tokens = min(tb.capacity, tb.tokens+refillTokens)
		tb.lastRefilledAt = time.Now()
		fmt.Printf("Bucket was refilled at %v with %v tokens\n", tb.lastRefilledAt.Local().Format("15:04:05"), tb.tokens)
	}

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}

	return false
}

func TokenBucketRateLimiter() {
	fmt.Println("Token Bucket Algorithm")
	tb := NewTokenBucket(10, 2)

	for i := range 15 {
		if tb.AllowRequest() {
			fmt.Printf("Request %d was allowed, current token count = %d\n", i, tb.tokens)
		} else {
			fmt.Printf("Request %d was denied, current token count = %d\n", i, tb.tokens)
		}
		time.Sleep(600 * time.Millisecond)
	}
}
