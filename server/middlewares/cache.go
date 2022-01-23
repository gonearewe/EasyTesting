package middlewares

import (
    "time"

    "github.com/patrickmn/go-cache"
)

// MemoryStore is a cache with a default expiration time of 2 hours, and which
// purges expired items every 30 minutes
var MemoryStore = cache.New(2*time.Hour, 30*time.Minute)