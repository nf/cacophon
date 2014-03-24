// +build appengine

package main

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"

	internal "google.golang.org/appengine/internal"
)

func main() {
	internal.Main()
}

func getCache(req *http.Request) (audio []byte) {
	c := appengine.NewContext(req)
	item, err := memcache.Get(c, cacheKey(req))
	if err != nil {
		if err != memcache.ErrCacheMiss {
			c.Errorf("getCache: %v", err)
		}
		return nil
	}
	return item.Value
}

func putCache(req *http.Request, audio []byte) {
	c := appengine.NewContext(req)
	if err := memcache.Set(c, &memcache.Item{
		Key:   cacheKey(req),
		Value: audio,
	}); err != nil {
		c.Errorf("putCache: %v", err)
	}
}

func cacheKey(req *http.Request) string {
	return req.URL.Query().Encode() // sorted by key
}
