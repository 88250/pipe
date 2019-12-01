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

package service

import (
	"testing"

	"github.com/88250/pipe/model"
)

func TestAddCategory(t *testing.T) {
	category := &model.Category{
		Title:           "测试添加的分类标题",
		Path:            "/category-test1",
		Description:     "测试描述",
		MetaKeywords:    "测试 Meta Keywords",
		MetaDescription: "测试 Meta Description",
		Tags:            "tag1,tag2",
		Number:          0,
		BlogID:          1,
	}

	if err := Category.AddCategory(category); nil != err {
		t.Errorf("add category failed: " + err.Error())

		return
	}
}

func TestConsoleGetCategories(t *testing.T) {
	categories, pagination := Category.ConsoleGetCategories(1, 1)
	if nil == categories {
		t.Errorf("categories is nil")

		return
	}
	if 1 != len(categories) {
		t.Errorf("expected is [%d], actual is [%d]", 1, len(categories))
	}
	if 1 != pagination.RecordCount {
		t.Errorf("expected is [%d], actual is [%d]", 1, pagination.RecordCount)
	}
}

func TestGetGetCategories(t *testing.T) {
	categories := Category.GetCategories(2, 1)
	if nil == categories {
		t.Errorf("categories is nil")
	}
}

func TestRemoveCategory(t *testing.T) {
	if err := Category.RemoveCategory(1, 1); nil != err {
		t.Errorf("remove category failed: " + err.Error())
	}
}
