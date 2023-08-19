package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	expires := time.Millisecond * 10
	cache := NewCache(expires)
	if cache.cache == nil {
		t.Error("Cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "keyz",
			inputVal: []byte("Value of the cache for keyz"),
		},
		{
			inputKey: "",
			inputVal: []byte("Value of the cache for "),
		},
		{
			inputKey: "key1",
			inputVal: []byte("Value of the cache for key1"),
		},
		{
			inputKey: "123",
			inputVal: []byte("Value of the cache for 123"),
		},
	}
	expires := time.Millisecond * 10
	cache := NewCache(expires)

	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)
		actual, ok := cache.Get(cas.inputKey)

		if !ok {
			t.Errorf("%s not found", cas.inputKey)
			continue
		}

		if string(actual) != string(cas.inputVal) {
			t.Errorf("%s does not match", cas.inputVal)
		}
	}
}

func TestReap(t *testing.T) {
	expires := time.Millisecond * 10
	cache := NewCache(expires)

	keyOne := "key1"
	cache.Add(keyOne, []byte("Value one"))
	time.Sleep(expires + time.Millisecond)

	_, ok := cache.Get(keyOne)

	if ok {
		t.Errorf("%s should have been deleted!", keyOne)
	}
}

func TestReapFail(t *testing.T) {
	expires := time.Millisecond * 10
	cache := NewCache(expires)

	keyOne := "key2"
	cache.Add(keyOne, []byte("Value one"))
	time.Sleep(expires / 2)

	_, ok := cache.Get(keyOne)

	if !ok {
		t.Errorf("%s should be present!", keyOne)
	}
}
