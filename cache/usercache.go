package cache

import (
	"github.com/b3log/pipe/model"
	"github.com/bluele/gcache"
)

var User = &userCache{
	idHolder: gcache.New(1024 * 10).LRU().Build(),
}

type userCache struct {
	idHolder gcache.Cache
}

func (cache *userCache) Put(user *model.User) {
	if err := cache.idHolder.Set(user.ID, user); nil != err {
		logger.Errorf("put user [id=%d] into cache failed: %s", user.ID, err)
	}
}

func (cache *userCache) Get(id uint) *model.User {
	ret, err := cache.idHolder.Get(id)
	if nil != err && gcache.KeyNotFoundError != err {
		logger.Errorf("get user [id=%d] from cache failed: %s", id, err)

		return nil
	}

	return ret.(*model.User)
}
