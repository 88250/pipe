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
	"database/sql"
	"os"
	"path/filepath"
	"testing"
)

func TestLegacySQLiteDatabase(t *testing.T) {
	// 复制由 go-sqlite3 v2.0.3 生成的数据库，避免修改兼容性样本。
	data, err := os.ReadFile("testdata/sqlite-v2.0.3.db")
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "legacy.db")
	if err = os.WriteFile(path, data, 0600); err != nil {
		t.Fatal(err)
	}
	database, err := sql.Open("sqlite3", path)
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()
	var integrity string
	if err = database.QueryRow("PRAGMA integrity_check").Scan(&integrity); err != nil || integrity != "ok" {
		t.Fatalf("database integrity: %q, error: %v", integrity, err)
	}
	var title, content string
	if err = database.QueryRow("SELECT title, content FROM legacy_articles WHERE id = 42").Scan(&title, &content); err != nil {
		t.Fatal(err)
	}
	if title != "Legacy article" || content != "历史正文：UTF-8 与 SQLite 数据兼容" {
		t.Fatalf("unexpected legacy article: %q, %q", title, content)
	}
	const updatedContent = "升级后的正文"
	if _, err = database.Exec("UPDATE legacy_articles SET content = ? WHERE id = 42", updatedContent); err != nil {
		t.Fatal(err)
	}
	if err = database.Close(); err != nil {
		t.Fatal(err)
	}
	reopened, err := sql.Open("sqlite3", path)
	if err != nil {
		t.Fatal(err)
	}
	defer reopened.Close()
	if err = reopened.QueryRow("SELECT content FROM legacy_articles WHERE id = 42").Scan(&content); err != nil {
		t.Fatal(err)
	}
	if content != updatedContent {
		t.Fatalf("updated content was not persisted: %q", content)
	}
}
