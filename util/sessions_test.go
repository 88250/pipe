// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package util

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
)

var sessionTestKey = []byte("pipe-upgrade-test-key-0123456789ab")

func legacySessionData() SessionData {
	return SessionData{
		UID: 42, UName: "legacy-user", UB3Key: "legacy-b3-key", URole: 2,
		UAvatar: "https://example.com/avatar.png", BID: 7, BURL: "https://example.com/blog",
	}
}

func TestSessionCookieRoundTrip(t *testing.T) {
	expected := legacySessionData()
	router := gin.New()
	store := cookie.NewStore(sessionTestKey)
	store.Options(sessions.Options{Path: "/", MaxAge: 3600, HttpOnly: true})
	router.Use(sessions.Sessions("pipe", store))
	router.POST("/login", func(c *gin.Context) {
		if err := expected.Save(c); err != nil {
			t.Error(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Status(http.StatusNoContent)
	})
	router.GET("/session", func(c *gin.Context) {
		c.JSON(http.StatusOK, GetSession(c))
	})
	response := httptest.NewRecorder()
	router.ServeHTTP(response, httptest.NewRequest(http.MethodPost, "/login", nil))
	if response.Code != http.StatusNoContent {
		t.Fatalf("login status: %d", response.Code)
	}
	cookies := response.Result().Cookies()
	if len(cookies) != 1 || cookies[0].Name != "pipe" || cookies[0].Path != "/" ||
		!cookies[0].HttpOnly || cookies[0].MaxAge != 3600 {
		t.Fatalf("unexpected session cookies: %+v", cookies)
	}
	for _, test := range []struct {
		name  string
		value string
		want  SessionData
	}{
		{"valid", cookies[0].Value, expected},
		{"tampered", cookies[0].Value + "!", SessionData{}},
	} {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/session", nil)
			request.AddCookie(&http.Cookie{Name: "pipe", Value: test.value})
			response := httptest.NewRecorder()
			router.ServeHTTP(response, request)
			var actual SessionData
			if err := json.Unmarshal(response.Body.Bytes(), &actual); err != nil {
				t.Fatal(err)
			}
			if response.Code != http.StatusOK || !reflect.DeepEqual(actual, test.want) {
				t.Fatalf("unexpected session: status=%d, data=%+v", response.Code, actual)
			}
		})
	}
}

func TestLegacySessionCookie(t *testing.T) {
	data, err := os.ReadFile("testdata/session-v0.0.5.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 固定样本由 gorilla/sessions v1.2.1 和 securecookie v1.1.1 生成，不校验样本时间。
	codec := securecookie.New(sessionTestKey, nil).MaxAge(0)
	values := map[interface{}]interface{}{}
	value := strings.TrimSpace(string(data))
	if err = codec.Decode("pipe", value, &values); err != nil {
		t.Fatal(err)
	}
	encoded, ok := values["data"].(string)
	if !ok {
		t.Fatal("legacy session data is not a string")
	}
	var actual SessionData
	if err = json.Unmarshal([]byte(encoded), &actual); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(actual, legacySessionData()) {
		t.Fatalf("unexpected legacy session: %+v", actual)
	}
	if err = codec.Decode("pipe", value+"!", &values); err == nil {
		t.Fatal("tampered legacy session was accepted")
	}
}
