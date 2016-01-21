package golly

import (
	"math"
	"time"
)

func RetryWithBackoff(n int, err error, t time.Duration) (time.Duration, error) {
	// n is number of errors, so n>=1
	// 2^3=8, so 2^5 = 32, so max backoff is about 30 s
	return time.Second * time.Duration(math.Pow(2, math.Min(float64(n), 5))), nil
}
