package golly

import (
	"time"
)

func RetryWithBackoff(n int, err error, time.Duration) (time.Duration, error) {
	// n is number of errors, so n>=1
	n = math.Min(float64(n), 5)
	// 2^3=8, so 2^5 = 32, so max backoff is about 30 s
	return time.Second * math.Pow(2,n)
}