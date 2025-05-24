package config

import "time"

type CacheConfig struct {
	CacheRedirectLinkTTL time.Duration `env:"CACHE_REDIRECT_LINK_TTL" default:"24h"`
}

var cacheConfig CacheConfig

func LoadCacheConfig() CacheConfig {
	if cacheConfig != (CacheConfig{}) {
		return cacheConfig
	}

	loadConfig(&cacheConfig)

	return cacheConfig
}
