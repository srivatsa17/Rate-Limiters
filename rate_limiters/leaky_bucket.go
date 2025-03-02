package rate_limiters

import (
	"fmt"
	"sync"
	"time"
)

// How does Leaky Bucket work ?
// 1. 	Consider a bucket having a capacity to hold certain number of tokens.
// 2.	There is a leak in the bucket at a constant rate
// 3.	Whenever the queue gets full, requests are denied, else allowed.

type LeakyBucket struct {
	capacity   int
	rate       time.Duration
	queue      chan struct{}
	mu         sync.Mutex
	leakTicker *time.Ticker
}

func NewLeakyBucket(capacity int, rate time.Duration) *LeakyBucket {
	lb := &LeakyBucket{
		capacity: capacity,
		rate:     rate,
		queue:    make(chan struct{}, capacity),
	}

	lb.leakTicker = time.NewTicker(rate)
	go lb.startLeaking()
	return lb
}

func (lb *LeakyBucket) startLeaking() {
	for range lb.leakTicker.C {
		lb.mu.Lock()
		if len(lb.queue) > 0 {
			<-lb.queue
		}
		lb.mu.Unlock()
	}
}

func (lb *LeakyBucket) AllowRequest() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	if len(lb.queue) < lb.capacity {
		lb.queue <- struct{}{}
		return true
	}

	return false
}

func LeakyBucketRateLimiter() {
	fmt.Println("Leaky Bucket Algorithm")
	lb := NewLeakyBucket(3, time.Second)

	for i := 1; i <= 10; i++ {
		if lb.AllowRequest() {
			fmt.Printf("%d Request allowed at %v.\n", i, time.Now().Local().Format("15:04:05"))
		} else {
			fmt.Printf("%d Request denied at %v.\n", i, time.Now().Local().Format("15:04:05"))
		}
		time.Sleep(600 * time.Millisecond)
	}
}
