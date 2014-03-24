// +build appengine

package main

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/memcache"

	internal "google.golang.org/appengine/internal"
)

func main() {
	internal.Main()
}

type Sound struct {
	Params []Param `datastore:",noindex"`
	Audio  []byte  `datastore:",noindex"`
}

func getCache(req *http.Request) (audio []byte) {
	c := appengine.NewContext(req)
	item, err := memcache.Get(c, cacheKey(req))
	if err == nil {
		return item.Value
	}
	// fallback to datastore
	if err != memcache.ErrCacheMiss {
		c.Errorf("getCache memcache error: %v", err)
	}
	var snd Sound
	err = datastore.Get(c, datastore.NewKey(c, "Sound", cacheKey(req), 0, nil), &snd)
	if err == nil {
		return snd.Audio
	}
	if err != datastore.ErrNoSuchEntity {
		c.Errorf("getCache datastore error: %v", err)
	}
	// new sound
	return nil
}

func putCache(req *http.Request, audio []byte) {
	c := appengine.NewContext(req)
	if err := memcache.Set(c, &memcache.Item{
		Key:   cacheKey(req),
		Value: audio,
	}); err != nil {
		c.Errorf("putCache memcache error: %v", err)
	}
	if _, err := datastore.Put(c, datastore.NewKey(c, "Sound", cacheKey(req), 0, nil), &Sound{
		Params: audioParam(req),
		Audio:  audio,
	}); err != nil {
		c.Errorf("putCache datastore error: %v", err)
	}
}

func cacheKey(req *http.Request) string {
	return req.URL.Query().Encode() // sorted by key
}

type Param struct {
	Key   string
	Value string
}

func audioParam(req *http.Request) (params []Param) {
	for k, v := range req.Form {
		params = append(params, Param{k, v[0]})
	}
	return
}
