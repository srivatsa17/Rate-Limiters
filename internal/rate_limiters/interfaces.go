package rate_limiters

type RateLimiter interface {
	AllowRequest() bool
}