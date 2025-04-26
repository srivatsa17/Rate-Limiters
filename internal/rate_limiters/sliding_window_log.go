package rate_limiters

import (
	"fmt"
	"sync"
	"time"
)

// How does Sliding Window Log work ?
// 1.	Keep a log of request timestamps.
// 2.	When a new request comes in, remove all entries older than the window size.
// 3.	Count the remaining entries.
// 4.	If the count is less than the limit, allow the request and add its timestamp to the log.
// 5.	If the count exceeds the limit, request is denied.

type SlidingWindowLog struct {
	capacity int
	window   time.Duration
	requests []time.Time
	mu       sync.Mutex
}

func NewSlidingWindowLog(capacity int, window time.Duration) *SlidingWindowLog {
	return &SlidingWindowLog{
		capacity: capacity,
		window:   window,
		requests: make([]time.Time, 0),
	}
}

func (sw *SlidingWindowLog) AllowRequest() bool {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	now := time.Now()
	sw.Cleanup(now)

	if len(sw.requests) < sw.capacity {
		sw.requests = append(sw.requests, now)
		return true
	}

	return false
}

func (sw *SlidingWindowLog) Cleanup(now time.Time) {
	expiration := now.Add(-sw.window)

	index := 0
	for i, request := range sw.requests {
		if request.After(expiration) {
			index = i
			break
		}
	}

	if index > 0 {
		fmt.Println("Cleanup done")
		sw.requests = sw.requests[index:]
	}
}
