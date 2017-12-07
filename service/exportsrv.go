// Pipe - A small and beautiful blogging platform written in golang.
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
	"regexp"

	"github.com/b3log/pipe/model"
)

var Export = &exportService{}

type exportService struct {
}

func (srv *exportService) ExportMarkdowns(blogID uint) (ret []*MarkdownFile) {
	var articles []*model.Article
	if err := db.Where("blog_id = ?", blogID).Find(&articles).Error; nil != err {
		logger.Errorf("export markdowns failed: " + err.Error())

		return
	}
	if 1 > len(articles) {
		return
	}

	for _, article := range articles {
		mdFile := &MarkdownFile{
			Name:    sanitizeFilename(article.Title),
			Content: article.Content,
		}

		ret = append(ret, mdFile)
	}

	return ret
}

func sanitizeFilename(unsanitized string) string {
	unsanitized = regexp.MustCompile("[\\?\\\\/:|<>\\*]").ReplaceAllString(unsanitized, " ") // filter out ? \ / : | < > *

	return regexp.MustCompile("\\s+").ReplaceAllString(unsanitized, "_") // white space as underscores
}
