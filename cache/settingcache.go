package cache

import (
	"fmt"
	"github.com/b3log/pipe/model"
	"github.com/bluele/gcache"
)

var Setting = &settingCache{
	categoryNameHolder: gcache.New(1024 * 10).LRU().Build(),
}

type settingCache struct {
	categoryNameHolder gcache.Cache
}

func (cache *settingCache) Put(setting *model.Setting) {
	if err := cache.categoryNameHolder.Set(fmt.Sprintf("%s-%s-%d", setting.Category, setting.Name, setting.BlogID), setting); nil != err {
		logger.Errorf("put setting [id=%d] into cache failed: %s", setting.ID, err)
	}
}

func (cache *settingCache) Get(category, name string, blogID uint) *model.Setting {
	ret, err := cache.categoryNameHolder.Get(fmt.Sprintf("%s-%s-%d", category, name, blogID))
	if nil != err {
		logger.Errorf("get setting [name=%s, category=%s, blogID=%d] from cache failed: %s", category, name, blogID, err)

		return nil
	}

	return ret.(*model.Setting)
}
