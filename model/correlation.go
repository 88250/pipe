// Solo.go - A small and beautiful golang blogging system, Solo's golang version.
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

package model

import "github.com/jinzhu/gorm"

// Correlation types.
const (
	CorrelationCategoryTag = iota
	CorrelationArticleTag
)

// Correlation model.
//   category_id - tag_id
//   article_id - tag_id
type Correlation struct {
	gorm.Model

	ID1  uint
	ID2  uint
	Type int `gorm:"size:16"` // 0: category-tag, 1: article-tag

	BlogID uint
}
