package golly

import (
	"time"
)

type Golly struct {
	panicH func(interface{}) error
	waitH  func(int, error, time.Duration) (time.Duration, error)

	failCount int
	failWait  time.Duration
}

// New returns a new golly struct appropriately configured.
func New() *Golly {
	return &Golly{
		waitH: RetryWithBackoff,
	}
}

func Panic(p func(interface{}) error) *Golly {
	return New().Panic(p)
}
func Retry(p func(int, error, time.Duration) (time.Duration, error)) *Golly {
	return New().Retry(p)
}
func Run(f func() error) error {
	return New().Run(f)
}

func (g *Golly) Panic(p func(interface{}) error) *Golly {
	g.panicH = p
	return g
}

func (g *Golly) Retry(wait func(count int, err error, dur time.Duration) (time.Duration, error)) *Golly {
	g.waitH = wait
	return g
}

func (g *Golly) Run(f func() error) (err error) {
	r := func() error {
		if nil != g.panicH {
			defer func() {
				if e := recover(); nil != e {
					err = g.panicH(e)
				}
			}()
		}
		return f()
	}
	for {
		err = r()
		if nil == err {
			g.failCount = 0
			return nil
		}
		g.failCount++
		if nil == g.waitH {
			return err
		}
		g.failWait, err = g.waitH(g.failCount, err, g.failWait)
		if nil != err {
			return err
		}
		if 0 < g.failWait {
			time.Sleep(g.failWait)
		}
	}
}
