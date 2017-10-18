// Solo.go - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017, b3log.org
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

package service

import (
	"sync"

	"github.com/b3log/solo.go/model"
	log "github.com/sirupsen/logrus"
)

var Tag = &tagService{
	mutex: &sync.Mutex{},
}

type tagService struct {
	mutex *sync.Mutex
}

func (srv *tagService) ConsoleGetTags(blogID uint) (ret []*model.Tag) {
	if err := db.Where(model.Tag{BlogID: blogID}).Order("article_count DESC, id DESC").Find(&ret).Error; nil != err {
		log.Errorf("get tags failed: " + err.Error())
	}

	return
}

func (srv *tagService) GetTags(size int, blogID uint) (ret []*model.Tag) {
	if err := db.Where(model.Tag{BlogID: blogID}).Order("article_count DESC, id DESC").Limit(size).Find(&ret).Error; nil != err {
		log.Errorf("get tags failed: " + err.Error())
	}

	return
}
