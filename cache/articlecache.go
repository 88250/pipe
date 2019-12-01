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

// Package cache includes caches.
package cache

import (
	"os"

	"github.com/88250/gulu"
	"github.com/88250/pipe/model"
	"github.com/bluele/gcache"
)

// Logger
var logger = gulu.Log.NewLogger(os.Stdout)

// Article cache.
var Article = &articleCache{
	idHolder: gcache.New(1024 * 10).LRU().Build(),
}

type articleCache struct {
	idHolder gcache.Cache
}

func (cache *articleCache) Put(article *model.Article) {
	if err := cache.idHolder.Set(article.ID, article); nil != err {
		logger.Errorf("put article [id=%d] into cache failed: %s", article.ID, err)
	}
}

func (cache *articleCache) Get(id uint) *model.Article {
	ret, err := cache.idHolder.Get(id)
	if nil != err && gcache.KeyNotFoundError != err {
		logger.Errorf("get article [id=%d] from cache failed: %s", id, err)

		return nil
	}
	if nil == ret {
		return nil
	}

	return ret.(*model.Article)
}
