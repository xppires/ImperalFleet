package interfaces

type RateLimiter interface { 
	Allow() bool
}