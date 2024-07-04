package cache

import (
	"testing"
	"time"
)

func TestCreateNilCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "key3",
			inputVal: []byte("val3"),
		},
	}

	for _, c := range cases {
		cache.Add(c.inputKey, c.inputVal)
		actual, ok := cache.Get(c.inputKey)

		if !ok {
			t.Errorf("%s not found", c.inputKey)
		}

		if string(actual) != string(c.inputVal) {
			t.Errorf("%s does not match %s", string(actual), string(c.inputVal))
		}
		continue
	}
}

func TestDelete(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(keyOne)

	if ok {
		t.Errorf("%s should have been deleted", keyOne)
	}
}
