package cache

import (
	"sync"
	"time"
)

type cacheEntry struct{
	createdAt time.Time
	val	  []byte
}

type Cache struct{
	Entries map[string]cacheEntry
	Mu	sync.Mutex
}
	
