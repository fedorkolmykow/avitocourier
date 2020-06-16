package limiter

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// CustomLimiter ...
type CustomLimiter interface {
    Limit(next http.Handler) http.Handler
}

type customLimiter struct{
	lim *rate.Limiter
}

func (c *customLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if c.lim.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// NewCustomLimiter returns a new CustomLimiter instance.
func NewCustomLimiter(burst int) CustomLimiter{
	lim := rate.NewLimiter(rate.Every(time.Minute), burst)
	return &customLimiter{lim: lim}
}
