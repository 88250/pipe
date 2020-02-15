// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package model

// Correlation types.
const (
	CorrelationCategoryTag = iota
	CorrelationArticleTag
	CorrelationBlogUser
	CorrelationArticleArchive
)

// Correlation model.
//   id1(category_id) - id2(tag_id)
//   id1(article_id) - id2(tag_id)
//   id1(blog_id) - id2(user_id) - int1(role) - int2(article_count)
//   id1(article_id) - id2(archive_id)
type Correlation struct {
	Model

	ID1  uint64 `json:"id1"`
	ID2  uint64 `json:"id2"`
	Str1 string `gorm:"size:255" json:"str1"`
	Str2 string `gorm:"size:255" json:"str2"`
	Str3 string `gorm:"size:255" json:"str3"`
	Str4 string `gorm:"size:255" json:"str4"`
	Int1 int    `json:"int1"`
	Int2 int    `json:"int2"`
	Int3 int    `json:"int3"`
	Int4 int    `json:"int4"`
	Type int    `json:"type"`

	BlogID uint64 `sql:"index" json:"blogID"`
}
