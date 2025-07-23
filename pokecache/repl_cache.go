package pokecache

import (
    //"fmt"
    "sync"
    "time"
)

type cacheEntry struct {
    createdAt time.Time
    val []byte
}

type Cache struct {
    data map[string]cacheEntry
    mu *sync.Mutex
}

var Interval time.Duration = 5 * time.Second
func NewCache(interval time.Duration) *Cache {
    newCache := &Cache{ data: map[string]cacheEntry{}, mu: &sync.Mutex{}}
    go newCache.reapLoop()
    return newCache
}

func (c *Cache) Add(key string, val []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = cacheEntry{createdAt: time.Now(), val: val}
    //fmt.Println(c.data)
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()
    cEntry, exists := c.data[key] 
    if exists {return cEntry.val, exists}
    return []byte(""), exists
}

func (c *Cache) reapLoop() {
    for {
        time.Sleep(1 * time.Second)
        for k, v := range c.data {
            createdAt := v.createdAt
            createdNw := time.Now()
            createdDf := createdNw.Sub(createdAt)        
            if createdDf > Interval {
                c.mu.Lock()
                delete(c.data, k)
                c.mu.Unlock()
            }
        }
    } 

}

