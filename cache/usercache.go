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
	"github.com/88250/pipe/model"
	"github.com/bluele/gcache"
)

// User cache.
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

func (cache *userCache) Get(id uint64) *model.User {
	ret, err := cache.idHolder.Get(id)
	if nil != err && gcache.KeyNotFoundError != err {
		logger.Errorf("get user [id=%d] from cache failed: %s", id, err)

		return nil
	}
	if nil == ret {
		return nil
	}

	return ret.(*model.User)
}
