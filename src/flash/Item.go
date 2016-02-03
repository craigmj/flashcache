package flash

import (
	"encoding/json"
	"time"
)

type Item struct {
	Key    string
	Expiry time.Time
	Raw    json.RawMessage
	Next   *Item
}

// NewItem creates a new cache Item with the given key and raw data,
// and adds the item into the cache, setting the last item pointer of the cache
func NewItem(leftNext **Item, key string, raw json.RawMessage, expires time.Duration) *Item {
	item := &Item{key, time.Now().Add(expires), raw, *leftNext}
	*leftNext = item
	return item
}

// Remove removes the item from the linked list. It requires the pointer to the
// element to its left, and returns the next item in the list.
func (item *Item) Remove(leftNext **Item) *Item {
	*leftNext = item.Next
	return item.Next
}

// Expired returns true if the item is expired
func (item *Item) Expired() bool {
	return time.Now().After(item.Expiry)
}
