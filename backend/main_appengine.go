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
	toCache <- cacheItem{req, audio}
}

type cacheItem struct {
	req   *http.Request
	audio []byte
}

var toCache = make(chan cacheItem)

func init() {
	http.HandleFunc("/_ah/start", startHandler)
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	for item := range toCache {
		go func(req *http.Request, audio []byte) {
			if err := memcache.Set(c, &memcache.Item{
				Key:   cacheKey(req),
				Value: audio,
			}); err != nil {
				c.Errorf("putCache: %v", err)
			}
		}(item.req, item.audio)
	}
}

func cacheKey(req *http.Request) string {
	return req.URL.Query().Encode() // sorted by key
}
