package utils

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// MemoryStore is a cache with a default expiration time of 2 hours, and which
// purges expired items every 1 minute
var MemoryStore = cache.New(2*time.Hour, 1*time.Minute)
