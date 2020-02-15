// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

// Package model is the "model" layer which defines entity structures with ORM and controller.
package model

// Archive model.
type Archive struct {
	Model

	Year         string `gorm:"size:4" json:"year"`
	Month        string `gorm:"size:2" json:"month"`
	ArticleCount int    `json:"articleCount"`

	BlogID uint64 `sql:"index" json:"blogID"`
}
