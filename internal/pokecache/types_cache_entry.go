package pokecache

import "time"

type CacheEntry struct {
	val       []byte
	createdAt time.Time
}
