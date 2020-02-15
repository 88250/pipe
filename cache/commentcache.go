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
	"time"

	"github.com/88250/pipe/model"
	"github.com/bluele/gcache"
)

// Comment service.
var Comment = &commentCache{
	idHolder: gcache.New(1024 * 10 * 10).LRU().Expiration(30 * time.Minute).Build(),
}

type commentCache struct {
	idHolder gcache.Cache
}

func (cache *commentCache) Put(comment *model.Comment) {
	if err := cache.idHolder.Set(comment.ID, comment); nil != err {
		logger.Errorf("put comment [id=%d] into cache failed: %s", comment.ID, err)
	}
}

func (cache *commentCache) Get(id uint) *model.Comment {
	ret, err := cache.idHolder.Get(id)
	if nil != err && gcache.KeyNotFoundError != err {
		logger.Errorf("get comment [id=%d] from cache failed: %s", id, err)

		return nil
	}
	if nil == ret {
		return nil
	}

	return ret.(*model.Comment)
}
