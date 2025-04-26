package fixedwindowcounter

import (
	"fmt"
	"sync"
	"time"
)

// How Fixed Window Counter works
// 1.	Time is divided into windows of fixed size
// 2.	Each window can hold/serve limited number of requests
// 3.	After the time/window has elapsed, the requests are reset

type FixedWindowCounter struct {
	capacity int
	window   time.Duration
	requests int
	mu       sync.Mutex
	timeStamp time.Time
}

func NewFixedWindowCounter(capacity int, window time.Duration) *FixedWindowCounter {
	return &FixedWindowCounter{
		capacity: capacity,
		window: window,
		timeStamp: time.Now(),
	}
}

func (fw *FixedWindowCounter) AllowRequest() bool {
	fw.mu.Lock()
	defer fw.mu.Unlock()

	// Check if the window has elapsed the current window, if yes, reset the timestamp and requests
	if time.Since(fw.timeStamp) > fw.window {
		fmt.Println("Resetting Fixed Window")
		fw.requests = 0
		fw.timeStamp = time.Now()
	}

	// Serve the requests
	if fw.requests < fw.capacity {
		fw.requests++
		return true
	}

	return false
}
