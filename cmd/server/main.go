package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/srivatsa17/Rate-Limiters/internal/rate_limiters"
	"github.com/srivatsa17/Rate-Limiters/internal/rate_limiters/fixedwindowcounter"
	"github.com/srivatsa17/Rate-Limiters/internal/rate_limiters/leakybucket"
	"github.com/srivatsa17/Rate-Limiters/internal/rate_limiters/slidingwindowcounter"
	"github.com/srivatsa17/Rate-Limiters/internal/rate_limiters/slidingwindowlog"
	"github.com/srivatsa17/Rate-Limiters/internal/rate_limiters/tokenbucket"
)

func main() {
	printMenu()

	choice, err := readUserChoice()
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number between 1 and 5.")
		os.Exit(1)
	}

	handleChoice(choice)
}

func printMenu() {
	fmt.Println("\nWelcome to Rate Limiting!")
	fmt.Println("Available Rate Limiter Algorithms:")
	fmt.Println("1. Token Bucket")
	fmt.Println("2. Leaky Bucket")
	fmt.Println("3. Fixed Window Counter")
	fmt.Println("4. Sliding Window Log")
	fmt.Println("5. Sliding Window Counter")
	fmt.Print("\nPlease enter the option number: ")
}

func readUserChoice() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input) // Remove all surrounding whitespace and newline
	return strconv.Atoi(input)
}

func handleChoice(choice int) {
	switch choice {
	case 1:
		fmt.Println("\nYou selected: Token Bucket Rate Limiter")
		tb := tokenbucket.NewTokenBucket(10, 2)
		rate_limiters.PerformRateLimiting(tb)
	case 2:
		fmt.Println("\nYou selected: Leaky Bucket Rate Limiter")
		lb := leakybucket.NewLeakyBucket(3, time.Second)
		rate_limiters.PerformRateLimiting(lb)
	case 3:
		fmt.Println("\nYou selected: Fixed Window Counter Rate Limiter")
		fw := fixedwindowcounter.NewFixedWindowCounter(3, time.Second)
		rate_limiters.PerformRateLimiting(fw)
	case 4:
		fmt.Println("\nYou selected: Sliding Window Log Rate Limiter")
		swl := slidingwindowlog.NewSlidingWindowLog(2, 2*time.Second)
		rate_limiters.PerformRateLimiting(swl)
	case 5:
		fmt.Println("\nYou selected: Sliding Window Counter Rate Limiter")
		swc := slidingwindowcounter.NewSlidingWindowCounter(3, 10*time.Second)
		rate_limiters.PerformRateLimiting(swc)
	default:
		fmt.Println("Invalid option. Please select a number between 1 and 5.")
		os.Exit(1)
	}
}