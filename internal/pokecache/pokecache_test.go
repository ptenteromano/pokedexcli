package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "test.com",
			val: []byte("testurl1"),
		},
		{
			key: "example.com",
			val: []byte("testurl2"),
		},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("Test case %v", tt.key), func(t *testing.T) {
			c := NewCache(5 * time.Second)
			c.Add(tt.key, tt.val)
			got, ok := c.Get(tt.key)

			if !ok {
				t.Errorf("Expected %v to be in the cache", tt.key)
			}

			if string(got) != string(tt.val) {
				t.Errorf("Expected %v, got %v", tt.val, got)
			}
		})
	}
}

func TestDropStale(t *testing.T) {
	c := NewCache(1 * time.Second)
	c.Add("test.com", []byte("testurl1"))

	_, ok := c.Get("test.com")

	if !ok {
		t.Errorf("Expected test.com to be in the cache")
		return
	}

	time.Sleep(2 * time.Second)

	_, ok = c.Get("test.com")

	if ok {
		t.Errorf("Expected test.com to be dropped from the cache")
	}
}
