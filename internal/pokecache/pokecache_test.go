package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
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
			inputVal: []byte("value1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("value2"),
		},
		{
			inputKey: "",
			inputVal: []byte("value3"),
		},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)
		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("%s not found", cas.inputKey)
			continue
		}

		if string(actual) != string(cas.inputVal) {
			t.Errorf("%s doesn't equal %s",
				string(actual),
				string(cas.inputVal),
			)
			continue
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	keyOne := "key1"
	cache.Add(keyOne, []byte("value1"))
	time.Sleep(interval + time.Millisecond)
	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("%s should have been reaped!", keyOne)
	}
}
