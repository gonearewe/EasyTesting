package middlewares

import (
    "time"

    "github.com/gin-gonic/gin"
    "github.com/ulule/limiter/v3"
    mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
    "github.com/ulule/limiter/v3/drivers/store/memory"
)

func rateLimit(maxRps int64) gin.HandlerFunc {
    // Define a limit rate to `maxRps` requests per second.
    rate := limiter.Rate{
        Period: 1 * time.Second,
        Limit:  maxRps,
    }
    // Create an in-memory store with a goroutine which clears expired keys.
    store := memory.NewStore()
    // Create the limiter instance which takes the store and the rate as arguments.
    instance := limiter.New(store, rate)
    // Create a new middleware with the limiter instance.
    middleware := mgin.NewMiddleware(instance)
    return middleware
}
