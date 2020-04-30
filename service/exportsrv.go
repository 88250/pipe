// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

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
	if err := db.Where("blog_id = ?", blogID).Find(&articles).Error; nil != err {
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
