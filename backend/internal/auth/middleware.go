package auth

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/httpx"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// RequireAdmin rejects requests without a valid bearer token.
func RequireAdmin(tokens *TokenManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := strings.CutPrefix(c.GetHeader("Authorization"), "Bearer ")
		if !ok || token == "" {
			httpx.Error(c, http.StatusUnauthorized, "missing bearer token")
			return
		}

		if _, err := tokens.Verify(token); err != nil {
			httpx.Error(c, http.StatusUnauthorized, "invalid or expired token")
			return
		}
		c.Next()
	}
}

// RateLimit allows each client IP `burst` requests at once, refilled at one
// request per `every`. It protects the login endpoint from brute force.
func RateLimit(every time.Duration, burst int) gin.HandlerFunc {
	limiters := newIPLimiters(rate.Every(every), burst)
	return func(c *gin.Context) {
		if !limiters.allow(c.ClientIP()) {
			httpx.Error(c, http.StatusTooManyRequests, "too many attempts, try again later")
			return
		}
		c.Next()
	}
}

// maxTrackedIPs bounds memory use; idle limiters are dropped past this size.
const maxTrackedIPs = 1024

type ipLimiters struct {
	mu       sync.Mutex
	limit    rate.Limit
	burst    int
	limiters map[string]*rate.Limiter
}

func newIPLimiters(limit rate.Limit, burst int) *ipLimiters {
	return &ipLimiters{limit: limit, burst: burst, limiters: make(map[string]*rate.Limiter)}
}

func (l *ipLimiters) allow(ip string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	limiter, ok := l.limiters[ip]
	if !ok {
		if len(l.limiters) >= maxTrackedIPs {
			l.dropIdle()
		}
		limiter = rate.NewLimiter(l.limit, l.burst)
		l.limiters[ip] = limiter
	}
	return limiter.Allow()
}

// dropIdle removes limiters that have refilled completely; forgetting them
// does not change what they would allow.
func (l *ipLimiters) dropIdle() {
	for ip, limiter := range l.limiters {
		if limiter.Tokens() >= float64(l.burst) {
			delete(l.limiters, ip)
		}
	}
}
