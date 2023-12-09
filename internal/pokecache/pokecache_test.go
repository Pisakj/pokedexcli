package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	interval := time.Second * 5
	cache := NewCache(interval)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		val      []byte
	}{
		{
			inputKey: "testKey",
			val:      []byte("testVal"),
		},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.val)
		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("cannot retrieve %v from cache", cas.val)
		}

		if string(actual) != string(cas.val) {
			t.Error("actual and expected values doesn't match")
		}
	}

}

func TestRemoveCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	testKey := "key1"
	cache.Add(testKey, []byte("val1"))
	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(testKey)

	if ok {
		t.Errorf("%s should be removed", testKey)
	}
}
