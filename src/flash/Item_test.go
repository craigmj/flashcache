package flash

import (
	"encoding/json"
	"testing"
	"time"
)

func TestItem(t *testing.T) {
	var first *Item
	raw := json.RawMessage([]byte("message"))
	item := NewItem(&first, "fred", raw, time.Minute)
	if item.Expired() {
		t.Errorf("Item just created but already expired")
	}
	if first != item {
		t.Errorf("First not set properly")
	}
	if nil != item.Next {
		t.Errorf("Item has non-nil Next pointer")
	}
	if "message" != string(item.Raw) {
		t.Errorf("Raw data on item not set properly")
	}

	second := NewItem(&first.Next, "joe", raw, time.Minute)
	if second.Expired() {
		t.Errorf("Second item just created by already expired")
	}
	if first.Next != second {
		t.Errorf("Created second item, but first.Next not set to second item")
	}
	if nil != second.Next {
		t.Errorf("Second item Next pointer is not nil")
	}

	if second != item.Remove(&first) {
		t.Errorf("Remove of single item not returning second")
	}
	if second != first {
		t.Errorf("Removed single item but first pointer not pointing to second")
	}
}
