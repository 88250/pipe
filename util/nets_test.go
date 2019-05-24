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

package util

import "testing"

func TestIsIP(t *testing.T) {
	if !IsIP("8.8.8.8") {
		t.Errorf("[8.8.8.8] is an IP")

		return
	}

	if IsIP("is.not.an.ip") {
		t.Errorf("[is.not.an.ip] is not an IP")

		return
	}

	if !IsIP("127.0.0.1") {
		t.Errorf("[127.0.0.1] is an IP")

		return
	}

	if IsIP("localhost") {
		t.Errorf("[localhost] is not an IP")

		return
	}
}

func TestIsDomain(t *testing.T) {
	if !IsDomain("b3log.org") {
		t.Errorf("[b3log.org] is a daomin")

		return
	}

	if IsDomain("localhost") {
		t.Errorf("[localhost] is not a domain")

		return
	}

	if IsDomain("8.8.8.8") {
		t.Errorf("[8.8.8.8] is not a domain")
	}
}

func TestIsBot(t *testing.T) {
	if !IsBot("Sym") {
		t.Errorf("[Sym] is not a bot")

		return
	}
}