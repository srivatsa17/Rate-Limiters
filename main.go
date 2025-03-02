package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/srivatsa17/Rate-Limiters/rate_limiters"
)

func main() {
	fmt.Println("Welcome to Rate Limiting\nPossible Options are:")
	fmt.Println("1. Token Bucket\n2. Leaky Bucket\n3. Fixed Window Counter\n4. Sliding Window Log\n5. Sliding Window Counter")
	fmt.Println("Please enter the option number of the Rate limiter algorithm you want to try out:")

	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	input = input[:len(input)-1] // Remove newline character
	choice, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		os.Exit(1)
	}

	switch choice {
	case 1:
		fmt.Println("You selected Token Bucket Rate Limiter")
		rate_limiters.TokenBucketRateLimiter()
	case 2:
		fmt.Println("You selected Leaky Bucket Rate Limiter")
		rate_limiters.LeakyBucketRateLimiter()
	case 3:
		fmt.Println("You selected Fixed Window Counter Rate Limiter")
		rate_limiters.FixedWindowCounterRateLimiter()
	case 4:
		fmt.Println("You selected Sliding Window Log Rate Limiter")
		// rate_limiters.SlidingWindowLogRateLimiter()
	case 5:
		fmt.Println("You selected Sliding Window Counter Rate Limiter")
		// rate_limiters.SlidingWindowCounterRateLimiter()
	default:
		fmt.Println("Invalid option. Please select a number between 1 and 5.")
		os.Exit(1)
	}
}
