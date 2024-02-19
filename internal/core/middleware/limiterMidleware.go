package middleware

import (
	"net/http"
	"time"

	"github.com/JubaerHossain/golang-ddd/internal/core/limiter"
	"github.com/JubaerHossain/golang-ddd/pkg/utils"
	"golang.org/x/time/rate"
)

func LimiterMiddleware(next http.Handler) http.Handler {
	var limiter = limiter.NewIPRateLimiter(rate.Every(time.Second), 1)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := limiter.GetLimiter(r.RemoteAddr)
		if !limiter.Allow() {
			utils.WriteJSONError(w, http.StatusTooManyRequests, "Too many requests")
			return
		}
		next.ServeHTTP(w, r)
	})
}
