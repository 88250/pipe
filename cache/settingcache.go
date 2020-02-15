// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package cache

import (
	"fmt"
	"time"

	"github.com/88250/pipe/model"
	"github.com/bluele/gcache"
)

// Setting cache.
var Setting = &settingCache{
	categoryNameHolder: gcache.New(1024 * 10).LRU().Expiration(30 * time.Minute).Build(),
}

type settingCache struct {
	categoryNameHolder gcache.Cache
}

func (cache *settingCache) Put(setting *model.Setting) {
	if err := cache.categoryNameHolder.Set(fmt.Sprintf("%s-%s-%d", setting.Category, setting.Name, setting.BlogID), setting); nil != err {
		logger.Errorf("put setting [id=%d] into cache failed: %s", setting.ID, err)
	}
}

func (cache *settingCache) Get(category, name string, blogID uint64) *model.Setting {
	ret, err := cache.categoryNameHolder.Get(fmt.Sprintf("%s-%s-%d", category, name, blogID))
	if nil != err && gcache.KeyNotFoundError != err {
		logger.Errorf("get setting [name=%s, category=%s, blogID=%d] from cache failed: %s", category, name, blogID, err)

		return nil
	}
	if nil == ret {
		return nil
	}

	return ret.(*model.Setting)
}
