package middlewares

import (
    "time"

    "github.com/gin-contrib/cache/persistence"
)

var MemoryStore = persistence.NewInMemoryStore(3*time.Hour)
