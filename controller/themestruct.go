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

package controller

type ThemeArticle struct {
	ID           uint
	Abstract     string
	Author       *ThemeAuthor
	CreatedAt    string
	Title        string
	Tags         []*ThemeTag
	URL          string
	Topped       bool
	ViewCount    int
	CommentCount int
	ThumbnailURL string
	Content      string
	Editable     bool
}

type ThemeTag struct {
	Title string
	URL   string
}

type ThemeAuthor struct {
	Name      string
	AvatarURL string
	URL       string
}

type ThemeCategory struct {
	Title string
	URL   string
}

type ThemeListComment struct {
	ID        uint
	Title     string
	Content   string
	URL       string
	Author    *ThemeAuthor
	CreatedAt string
	Removable bool
}
