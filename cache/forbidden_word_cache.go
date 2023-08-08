package cache

import (
	"time"
	"vbbs/model"
	"vbbs/repositories"

	"github.com/goburrow/cache"
	"github.com/sirupsen/logrus"
	"vbbs/pkg/simple/sqls"
)

type forbiddenWordCache struct {
	cache cache.LoadingCache
}

var ForbiddenWordCache = newForbiddenWordCache()

func newForbiddenWordCache() *forbiddenWordCache {
	return &forbiddenWordCache{
		cache: cache.NewLoadingCache(
			func(_ cache.Key) (value cache.Value, e error) {
				value = repositories.ForbiddenWordRepository.Find(sqls.DB(), sqls.NewCnd())
				return
			},
			cache.WithMaximumSize(1000),
			cache.WithExpireAfterAccess(30*time.Minute),
		),
	}
}

func (c *forbiddenWordCache) Get() []model.ForbiddenWord {
	val, err := c.cache.Get("_")
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return val.([]model.ForbiddenWord)
}

func (c *forbiddenWordCache) Invalidate() {
	c.cache.Invalidate("_")
}
