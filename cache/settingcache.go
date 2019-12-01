// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cache

import (
	"fmt"

	"github.com/88250/pipe/model"
	"github.com/bluele/gcache"
)

// Setting cache.
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
