package stm

import (
	"fmt"
	"net/http"
	"time"
)

// PingSearchEngines requests some ping server from it calls Sitemap.PingSearchEngines.
func PingSearchEngines(opts *Options, urls ...string) {
	urls = append(urls, []string{
		"http://www.google.com/webmasters/tools/ping?sitemap=%s",
		"http://www.bing.com/webmaster/ping.aspx?siteMap=%s",
	}...)
	sitemapURL := opts.IndexLocation().URL()

	bufs := len(urls)
	does := make(chan string, bufs)
	client := http.Client{Timeout: time.Duration(5 * time.Second)}

	for _, url := range urls {
		go func(baseurl string) {
			url := fmt.Sprintf(baseurl, sitemapURL)
			println("Ping now:", url)

			resp, err := client.Get(url)
			if err != nil {
				does <- fmt.Sprintf("[E] Ping failed: %s (URL:%s)",
					err, url)
				return
			}
			defer resp.Body.Close()

			does <- fmt.Sprintf("Successful ping of `%s`", url)
		}(url)
	}

	for i := 0; i < bufs; i++ {
		println(<-does)
	}
}
