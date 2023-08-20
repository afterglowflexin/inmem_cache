package main

import (
	"fmt"
	"sync"
	"time"
)

// ВОПРОСЫ
// 1. Использование поинтеров в profiles и dieTimes в кэше (на саму map'у и объекты внутри map'ы)
func main() {

}

type Cache struct {
	*sync.Mutex
	profiles map[string]Profile
	dieTimes map[string]time.Time
	doneChan chan int
	*sync.Once
}

func NewCache() *Cache {
	c := &Cache{
		Mutex:    &sync.Mutex{},
		profiles: make(map[string]Profile),
		dieTimes: make(map[string]time.Time),
	}
	go c.Clear(50 * time.Millisecond)
	return c
}

func (c *Cache) Clear(delay time.Duration) {
	for {
		select {
		case <-time.After(delay):
			for id, t := range c.dieTimes {
				if time.Now().Sub(t) >= 0 {
					delete(c.dieTimes, id)
					delete(c.profiles, id)
				}
			}
		case <-c.doneChan:
			return
		}
	}
}

func (c *Cache) Set(profile Profile) {
	c.Lock()
	defer c.Unlock()
	c.profiles[profile.UUID] = profile
	c.dieTimes[profile.UUID] = time.Now().Add(2 * time.Second)
}

func (c *Cache) Get(uuid string) (Profile, error) {
	if profile, ok := c.profiles[uuid]; ok || c.dieTimes[uuid].Sub(time.Now()) > 0 {
		return profile, nil
	}
	return Profile{}, fmt.Errorf("no such profile")
}

func (c *Cache) Close() {
	c.Once.Do(func() {
		close(c.doneChan)
	})
}

type Profile struct {
	UUID   string
	Name   string
	Orders []*Order
}

type Order struct {
	UUID      string
	Value     any
	CreatedAt time.Time
	UpdatedAt time.Time
}
