package ratelimiter

import (
	"fmt"
	"time"

	"github.com/Kartik-Kumar12/Rate-Limiter-Algorithms/rate_limiter/algorithms"
)

func GetRateLimiter(algo string) (Ratelimiter, error) {
	switch algo {
	case "token":
		return algorithms.NewTokenBucket(5, 1), nil
	case "leaky":
		return algorithms.NewLeakyBucket(5, 1), nil
	case "fixed_window":
		return algorithms.NewWindow(4, 1*time.Second), nil
	case "sliding_window_log":
		return algorithms.NewSlidingWindowLog(4, 1*time.Second), nil
	case "sliding_window_counter":
		return algorithms.NewSlidingWindowCounter(5, 5, 1*time.Second), nil
	default:
		return nil, fmt.Errorf("unsupported algorithm %v", algo)
	}
}
