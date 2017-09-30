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
	"math"
	"sync"

	"github.com/b3log/solo.go/model"
	"github.com/b3log/solo.go/util"
)

var Navigation = &navigationService{
	mutex: &sync.Mutex{},
}

type navigationService struct {
	mutex *sync.Mutex
}

// Navigation pagination arguments of admin console.
const (
	adminConsoleNavigationListPageSize    = 15
	adminConsoleNavigationListWindowsSize = 20
)

func (srv *navigationService) ConsoleGetNavigations(page int, blogID uint) (ret []*model.Navigation, pagination *util.Pagination) {
	offset := (page - 1) * adminConsoleNavigationListPageSize
	count := 0
	db.Model(model.Navigation{}).Order("number ASC, id DESC").
		Where(model.Navigation{BlogID: blogID}).
		Count(&count).Offset(offset).Limit(adminConsoleNavigationListPageSize).Find(&ret)

	pageCount := int(math.Ceil(float64(count) / adminConsoleNavigationListPageSize))
	pagination = util.NewPagination(page, adminConsoleNavigationListPageSize, pageCount, adminConsoleNavigationListWindowsSize, count)

	return
}

func (srv *navigationService) ConsoleGetNavigation(id uint) *model.Navigation {
	ret := &model.Navigation{}
	if nil != db.First(ret, id).Error {
		return nil
	}

	return ret
}
