package shared

import (
	"context"
	"time"
)

type RateLimiter struct {
	ticker *time.Ticker
	done   chan struct{}
}

func NewRateLimiter(delay time.Duration, rps int) *RateLimiter {
	if rps > 0 {
		delay = time.Second / time.Duration(rps)
	}

	if delay <= 0 {
		return nil
	}

	limiter := &RateLimiter{
		ticker: time.NewTicker(delay),
		done:   make(chan struct{}),
	}

	return limiter
}

func (r *RateLimiter) Wait(ctx context.Context) error {
	if r == nil {
		return nil
	}

	select {
	case <-r.ticker.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (r *RateLimiter) Stop() {
	if r != nil {
		r.ticker.Stop()
		close(r.done)
	}
}
