package golly

import (
	"fmt"
	"time"
)

type Golly struct {
	panicH func(interface{}) error
	waitH  func(n int, func(int, error, time.Duration) (time.Duration,error)

	failCount int
	failWait time.Duration
}

// New returns a new golly struct appropriately configured.
func New() *Golly {
	return &Golly{}
}

func (g *Golly) Panic(p func(interface{}) error) *Golly {
	g.panicH = p
	return g
}

func (g *Golly) Retry(wait func (count int, err error, dur time.Duration) (time.Duration, error)) *Golly {
	g.waitH = wait
	return g
}

func (g *Golly) Run(f func() error) (err error) {
	r := func() error {
		if nil != g.panicH {
			defer func() {
				if e:=recover(); nil!=e {
					err = g.panicH(e)
				}
			}()
		}
		return f()
	}
	for {
		err = r()
		if nil==err {
			g.failCount = 0
			return nil
		}
		g.failCount++
		if nil!=g.waitH {
			wait, err := g.waitH(g.failCount, err, g.failWait)
			if nil!=err {
				return err
			}
			if 0<wait {
				time.Sleep(wait)
			}
		}
	}
}

