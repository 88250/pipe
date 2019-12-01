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
	"regexp"
	"strings"

	"github.com/88250/pipe/model"
	"gopkg.in/yaml.v2"
)

// Export service.
var Export = &exportService{}

type exportService struct {
}

func (srv *exportService) ExportMarkdowns(blogID uint64) (ret []*MarkdownFile) {
	var articles []*model.Article
	if err := db.Where("`blog_id` = ?", blogID).Find(&articles).Error; nil != err {
		logger.Errorf("export markdowns failed: " + err.Error())

		return
	}
	if 1 > len(articles) {
		return
	}

	for _, article := range articles {
		front := struct {
			Title     string   `yaml:"title"`
			Date      string   `yaml:"date"`
			Updated   string   `yaml:"updated"`
			Tags      []string `yaml:"tags"`
			Permalink string   `yaml:"permalink"`
		}{
			article.Title,
			article.CreatedAt.Format("2006-01-02 15:04:05"),
			article.UpdatedAt.Format("2006-01-02 15:04:05"),
			strings.Split(article.Tags, ","),
			article.Path,
		}
		frontData, err := yaml.Marshal(front)
		if nil != err {
			logger.Errorf("marshal front matter failed: " + err.Error())

			continue
		}

		mdFile := &MarkdownFile{
			Name:    sanitizeFilename(article.Title),
			Content: string(frontData) + "---\n" + article.Content,
		}

		ret = append(ret, mdFile)
	}

	return ret
}

func sanitizeFilename(unsanitized string) string {
	unsanitized = regexp.MustCompile("[\\?\\\\/:|<>\\*]").ReplaceAllString(unsanitized, " ") // filter out ? \ / : | < > *

	return regexp.MustCompile("\\s+").ReplaceAllString(unsanitized, "_") // white space as underscores
}
