package repository

import (
	"context"
	"go-boilerplate/internal/books/domain/books"
	"go-boilerplate/shared/utils"

	"github.com/redis/go-redis/v9"
	"github.com/ztrue/tracerr"
)

type RedisBooksCache struct {
	cache *redis.Client
}

func NewBooksRedisCache(cache *redis.Client) Cache {
	return &RedisBooksCache{cache}
}

func (r RedisBooksCache) GetBooks(c context.Context, request books.GetBooksDto) ([]books.BookWithAuthor, error) {
	title := utils.ValueOrEmptyString(request.Title)
	key := utils.NewRedisKey("books", title)

	result, err := r.cache.Get(c, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}

		return nil, tracerr.Wrap(err)
	}

	var jsonBooks []books.BookWithAuthor
	err = utils.FromJsonString(result, &jsonBooks)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return jsonBooks, nil
}

func (r RedisBooksCache) SetBooks(c context.Context, request books.GetBooksDto, books []books.BookWithAuthor) error {
	title := utils.ValueOrEmptyString(request.Title)
	key := utils.NewRedisKey("books", title)

	value, err := utils.ToJsonString(books)
	if err != nil {
		return tracerr.Wrap(err)
	}

	err = r.cache.Set(c, key, value, 0).Err()
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}

func (r RedisBooksCache) ClearBooks(c context.Context) error {
	key := utils.NewRedisKey("books", "*")

	iter := r.cache.Scan(c, 0, key, 0).Iterator()
	for iter.Next(c) {
		if err := iter.Err(); err != nil {
			return tracerr.Wrap(err)
		}

		err := r.cache.Del(c, iter.Val()).Err()
		if err != nil {
			return tracerr.Wrap(err)
		}
	}

	return nil
}
