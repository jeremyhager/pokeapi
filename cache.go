package pokeapi

import "time"

var CacheSettings = Settings{
	CustomExpire: 0,
	UseCache:     true,
}

var defaultCacheSettings = defaultSettings{
	MaxExpire: 10 * time.Minute,
	MinExpire: 5 * time.Minute,
}

// defaultSettings are settings not meant to be changed.
type defaultSettings struct {
	MaxExpire time.Duration
	MinExpire time.Duration
}

// CacheSettings are user settings for cache expiration.
type Settings struct {
	CustomExpire time.Duration
	UseCache     bool
}

// ClearCache clears all cached data.
func (c *RESTClient) ClearCache() {
	c.cache.Flush()
}

// setCache adds new item to local cache.
func (c *RESTClient) setCache(endpoint string, body []byte) {
	if CacheSettings.CustomExpire != 0 {
		c.cache.Set(endpoint, body, CacheSettings.CustomExpire*time.Minute)
	} else {
		c.cache.SetDefault(endpoint, body)
	}
}
