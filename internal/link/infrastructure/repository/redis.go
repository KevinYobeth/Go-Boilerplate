package repository

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/config"
	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/utils"
	"github.com/redis/go-redis/v9"
	"github.com/ztrue/tracerr"
)

type RedisLinkCache struct {
	cache *redis.Client
}

func NewLinkRedisCache(cache *redis.Client) Cache {
	return &RedisLinkCache{cache}
}

func (r *RedisLinkCache) GetRedirectLink(c context.Context, slug string) (*link.RedirectLink, error) {
	key := RedirectLinkKey(slug)

	result, err := r.cache.Get(c, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}

		return nil, tracerr.Wrap(err)
	}

	var redirectLink *link.RedirectLink
	err = utils.FromJsonString(result, &redirectLink)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return redirectLink, nil
}

func (r *RedisLinkCache) SetRedirectLink(c context.Context, slug string, value link.RedirectLink) error {
	key := RedirectLinkKey(slug)
	config := config.LoadCacheConfig()

	redirectLink, err := utils.ToJsonString(value)
	if err != nil {
		return tracerr.Wrap(err)
	}

	err = r.cache.Set(c, key, redirectLink, config.CacheRedirectLinkTTL).Err()
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}

func (r *RedisLinkCache) ClearRedirectLink(c context.Context, slug string) error {
	key := RedirectLinkKey(slug)

	err := r.cache.Del(c, key).Err()
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}
