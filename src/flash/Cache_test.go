package flash

import (
	"encoding/json"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c := &Cache{}

	c.Add("one", json.RawMessage([]byte("one.1")), time.Minute)
	if nil == c.First {
		t.Errorf("Added a value, but First not set")
	}
	if c.First != c.Last {
		t.Errorf("Only one value in cache, but first!=last")
	}
	if nil != c.First.Next {
		t.Errorf("One value in cache, but First.Next != nil")
	}

	c.Add("one", json.RawMessage([]byte("one.2")), time.Minute)
	if c.First == c.Last {
		t.Errorf("Two values in cache, but first==last")
	}
	if c.First.Next != c.Last {
		t.Errorf("Two values in cache, but First.Next != Last (First.Next = %p)", c.First.Next)
	}

	two := c.Find("two")
	if 0 < len(two) {
		t.Errorf("Found %d values for key two, but key two hadn't been set", len(two))
	}

	one := c.Find("one")
	if 2 != len(one) {
		t.Errorf("Found %d values for key one, but 2 expected", len(one))
	}

	if nil != c.First {
		t.Errorf("Found items for key one - cache should be clear, but c.First still set")
	}
	if nil != c.Last {
		t.Errorf("Found items for key one - cache should be clear, but c.Last still set")
	}
}
