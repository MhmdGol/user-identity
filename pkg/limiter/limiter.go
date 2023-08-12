package limiter

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type IPLimiter struct {
	limiterMu sync.Mutex
	ipLimiter map[string]*ipLimit
}

func NewIPLimiter(cleanupTimer time.Duration) *IPLimiter {
	l := &IPLimiter{
		limiterMu: sync.Mutex{},
		ipLimiter: make(map[string]*ipLimit),
	}

	go func() {
		for {
			l.CleanupIPlimiter(cleanupTimer)
			time.Sleep(cleanupTimer)
		}
	}()

	return l
}

type ipLimit struct {
	limiter *rate.Limiter
	lastHit time.Time
}

func (il *IPLimiter) Limiter(ip string) *rate.Limiter {
	il.limiterMu.Lock()
	defer il.limiterMu.Unlock()

	iplim, exists := il.ipLimiter[ip]
	if !exists {
		limiter := rate.NewLimiter(rate.Every(time.Second*2), 5)
		iplim = &ipLimit{
			limiter: limiter,
			lastHit: time.Now(),
		}
		il.ipLimiter[ip] = iplim
	}

	iplim.lastHit = time.Now()
	return iplim.limiter
}

func (il *IPLimiter) CleanupIPlimiter(inactivateDuration time.Duration) {
	il.limiterMu.Lock()
	defer il.limiterMu.Unlock()

	for ip, iplim := range il.ipLimiter {
		if time.Since(iplim.lastHit) > inactivateDuration {
			delete(il.ipLimiter, ip)
		}
	}
}
