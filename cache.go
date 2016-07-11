// Package couchcache provides an implementation of httpcache.Cache that stores and
// retrieves data using Memcached.
package memcached

import (
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

// Cache objects store and retrieve data using Memcached.
type Cache struct {
	// couch database where the cache will be stored
	client *memcache.Client
}

// New returns a new Cache
func New(c *memcache.Client) *Cache {
	return &Cache{
		client: c,
	}
}

func (c *Cache) Get(key string) (resp []byte, ok bool) {
	i, err := c.client.Get(key)
	if err != nil {
		return []byte{}, false
	}
	return i.Value, true
}

func (c *Cache) Set(key string, content []byte) {
	err := c.client.Set(&memcache.Item{
		Key:   key,
		Value: content,
	})
	if err != nil {
		log.Printf("Can't insert record in memcache: %v\n", err)
	}
	return
}

func (c *Cache) Delete(key string) {
	err := c.client.Delete(key)
	if err != nil {
		log.Printf("Can't remove record from memcache %s", err)
	}
}

func (c *Cache) Indexes() {}
