package rate_limiters

import (
	"fmt"
	"time"
)

func PerformRateLimiting(rl RateLimiter) {
	for i := range 15 {
		if rl.AllowRequest() {
			fmt.Printf("Request %d was allowed\n", i)
		} else {
			fmt.Printf("Request %d was denied\n", i)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
