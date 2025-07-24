package middleware

import (
    "app/internal/interfaces"
    "app/internal/common"
	"net/http"  
)

type GobalRateMiddleWare  struct {
     limiter  interfaces.RateLimiter
} 

func NewGobalRateMiddleWare(limiter interfaces.RateLimiter) *GobalRateMiddleWare {
    return &GobalRateMiddleWare{limiter:limiter} 
}

func (gr *GobalRateMiddleWare ) RateLimitMiddleware(next http.Handler ) http.Handler {
 
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !gr.limiter.Allow() { 
            common.HandleError(w, nil, http.StatusTooManyRequests, "Rate limit exceeded")		
            return
        }
        next.ServeHTTP(w, r)
    })
}