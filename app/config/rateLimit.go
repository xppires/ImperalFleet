package config

import (
	"golang.org/x/time/rate"
	"app/internal/interfaces"
	"os"
	"strconv"
)
 
func  InitGlobalLimitRate() interfaces.RateLimiter {
	 
	limit, err := strconv.Atoi(os.Getenv("GLOBAL_RATE"))
    if err != nil {
        // ... handle error
        panic(err)
    }
	brusts, err := strconv.Atoi(os.Getenv("GLOBAL_RATE_BURSTS"))
	 if err != nil {
        // ... handle error
        panic(err)
    }
	l :=  rate.NewLimiter(rate.Limit(limit), brusts)
	var gl interfaces.RateLimiter 

	gl = l
 
  return  gl
}