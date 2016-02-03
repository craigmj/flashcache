package flash

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Cache struct {
	Lock  sync.Mutex
	First *Item
	Last  *Item
}

// Add adds the raw value to the cache for the given key
func (c *Cache) Add(key string, raw json.RawMessage, expires time.Duration) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	// if the Last is nil, then the First will also be nil
	if nil == c.Last {
		NewItem(&c.Last, key, raw, expires)
		c.First = c.Last
	} else {
		NewItem(&c.Last.Next, key, raw, expires)
		c.Last = c.Last.Next
	}
}

// FlushExpired removes all expired items from the cache
func (c *Cache) FlushExpired() {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	for nil != c.First && c.First.Expired() {
		c.First = c.First.Remove(&c.First)
	}
	if nil == c.First {
		c.Last = nil
	}
}

func (c *Cache) Find(key string) []json.RawMessage {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	left := &c.First
	item := c.First
	raws := make([]json.RawMessage, 0)

	for nil != item {
		if item.Expired() {
			item = item.Remove(left)
			continue
		}
		if item.Key == key {
			raws = append(raws, item.Raw)
			item = item.Remove(left)
			continue
		}
		left = &(item.Next)
		item = *left
	}
	if nil == c.First {
		c.Last = nil
	}
	return raws
}

func (c *Cache) WebDump(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	for i := c.First; nil != i; i = i.Next {
		fmt.Fprintln(w, "%s : %s\n", i.Key, string(i.Raw))
	}
}

func (c *Cache) Len() int {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	i := 0
	for item := c.First; nil != item; item = item.Next {
		i++
	}
	return i
}
