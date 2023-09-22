package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var CacheInstance *cache.Cache

func init() {
	CacheInstance = cache.New(24*time.Hour, 24*time.Hour)
}
