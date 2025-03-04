package rate_limiters

import (
	"fmt"
	"sync"
	"time"
)

// How does Sliding Window Counter work ?
// 1.	Keep track of request count for the current and previous window.
// 2.	Calculate the weighted sum of requests based on the overlap with the sliding window.
// 3.	If the weighted sum is less than the limit, allow the request.

type SlidingWindowCounter struct {
	capacity      int
	window        time.Duration
	currentCount  int
	previousCount int
	mu            sync.Mutex
	lastTick      time.Time
}

func NewSlidingWindowCounter(capacity int, window time.Duration) *SlidingWindowCounter {
	return &SlidingWindowCounter{
		capacity: capacity,
		window:   window,
		lastTick: time.Now(),
	}
}

func (sw *SlidingWindowCounter) AllowRequest() bool {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	now := time.Now()
	elapsed := time.Since(sw.lastTick)

	if elapsed >= sw.window {
		sw.previousCount = sw.currentCount
		sw.currentCount = 0
		sw.lastTick = now
	} else if elapsed > 0 {
		weight := float64(sw.window - elapsed) / float64(sw.window)
		threshold := sw.previousCount * int(weight) + sw.currentCount

		if threshold >= sw.capacity {
			return false
		}
	}

	sw.currentCount++
	return true
}

func SlidingWindowCounterRateLimiter() {
	fmt.Println("Rate Limiter with Sliding Window Counter")
	sw := NewSlidingWindowCounter(3, 10 * time.Second)

	for i := 1; i <= 10; i++ {
		if sw.AllowRequest() {
			fmt.Printf("%d Request allowed at %v.\n", i, time.Now().Local().Format("15:04:05"))
		} else {
			fmt.Printf("%d Request allowed at %v.\n", i, time.Now().Local().Format("15:04:05"))
		}
		time.Sleep(2 * time.Second)
	}
}
