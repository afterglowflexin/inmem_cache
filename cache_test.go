package main

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	mockProfile := Profile{
		UUID: "abc-def-poi",
		Name: "Andrew",
		Orders: []*Order{
			&Order{
				UUID:      "abc-def-poi",
				Value:     1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}

	cache := NewCache()
	cache.Add(mockProfile)
	time.Sleep(3 * time.Second)
	_, err := cache.Get(mockProfile.UUID)
	if err == nil {
		t.Errorf("want err no such profile, got nil")
	}
}
