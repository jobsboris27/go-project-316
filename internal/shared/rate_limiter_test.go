package shared

import (
	"context"
	"testing"
	"time"
)

func TestRateLimiter_Delay(t *testing.T) {
	limiter := NewRateLimiter(50*time.Millisecond, 0)
	defer limiter.Stop()

	ctx := context.Background()
	start := time.Now()

	_ = limiter.Wait(ctx)
	_ = limiter.Wait(ctx)
	_ = limiter.Wait(ctx)

	elapsed := time.Since(start)

	if elapsed < 100*time.Millisecond {
		t.Errorf("Expected at least 100ms for 3 waits with 50ms delay, got %v", elapsed)
	}
}

func TestRateLimiter_RPS(t *testing.T) {
	limiter := NewRateLimiter(0, 10)
	defer limiter.Stop()

	ctx := context.Background()
	start := time.Now()

	for i := 0; i < 5; i++ {
		_ = limiter.Wait(ctx)
	}

	elapsed := time.Since(start)

	if elapsed < 400*time.Millisecond {
		t.Errorf("Expected at least 400ms for 5 requests at 10 RPS, got %v", elapsed)
	}
}

func TestRateLimiter_NoLimit(t *testing.T) {
	limiter := NewRateLimiter(0, 0)

	if limiter != nil {
		t.Error("RateLimiter should be nil when no limits set")
	}
}

func TestRateLimiter_ContextCancellation(t *testing.T) {
	limiter := NewRateLimiter(1*time.Second, 0)
	defer limiter.Stop()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := limiter.Wait(ctx)

	if err == nil {
		t.Error("Wait should return error when context is cancelled")
	}
}

func TestRateLimiter_RPSPriority(t *testing.T) {
	limiter := NewRateLimiter(1*time.Millisecond, 100)
	defer limiter.Stop()

	ctx := context.Background()
	start := time.Now()

	for i := 0; i < 3; i++ {
		_ = limiter.Wait(ctx)
	}

	elapsed := time.Since(start)

	if elapsed < 20*time.Millisecond {
		t.Errorf("RPS should take priority, expected ~20ms, got %v", elapsed)
	}
}
