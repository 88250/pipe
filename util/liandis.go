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
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/88250/gulu"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

// Logger
var logger = gulu.Log.NewLogger(os.Stdout)

// CommunityURL is the URL of LianDi community.
const CommunityURL = "https://ld246.com"

// CommunityAPI is a reverse proxy for https://ld246.com.
func CommunityAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		proxy := httputil.NewSingleHostReverseProxy(&url.URL{
			Scheme: "https",
			Host:   "ld246.com",
		})

		proxy.Transport = &http.Transport{DialTLS: dialTLS}
		director := proxy.Director
		proxy.Director = func(req *http.Request) {
			director(req)
			req.Host = req.URL.Host
			req.URL.Path = req.URL.Path[len("api/hp/"):]
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func dialTLS(network, addr string) (net.Conn, error) {
	conn, err := net.Dial(network, addr)
	if err != nil {
		return nil, err
	}

	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	cfg := &tls.Config{ServerName: host}

	tlsConn := tls.Client(conn, cfg)
	if err := tlsConn.Handshake(); err != nil {
		conn.Close()
		return nil, err
	}

	cs := tlsConn.ConnectionState()
	cert := cs.PeerCertificates[0]

	cert.VerifyHostname(host)

	return tlsConn, nil
}

// HacPaiUserInfo returns HacPai community user info specified by the given access token.
func HacPaiUserInfo(accessToken string) (ret map[string]interface{}) {
	result := map[string]interface{}{}
	response, data, errors := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		Post(CommunityURL+"/user/ak").SendString("access_token="+accessToken).Timeout(7*time.Second).
		Set("User-Agent", UserAgent).EndStruct(&result)
	if nil != errors || http.StatusOK != response.StatusCode {
		logger.Errorf("get community user info failed: %+v, %s", errors, data)
		return nil
	}
	if 0 != result["code"].(float64) {
		return nil
	}
	return result["data"].(map[string]interface{})
}
