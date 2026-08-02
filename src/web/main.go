package web

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)
import (
	"bypaws/src/proxy"
)

const targetSite = "https://e621.net/"

var allowedPaths = []string{
	"/favicon.ico",
	"/vite",
	"/images",
	"/manifest.json",

	"/posts",
	"/popular",
}

func Run() {
	targetURL, _ := url.Parse(targetSite)

	proxyURL, err := proxy.GetFirstProxy()
	if err != nil {
		log.Fatalf("프록시 로드 실패: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Transport = &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		isAllowed := (req.URL.Path == "/")

		if !isAllowed {
			for _, path := range allowedPaths {
				if strings.HasPrefix(req.URL.Path, path) {
					isAllowed = true
					break
				}
			}
		}

		if !isAllowed {
			http.Error(
				writer,
				"해당 페이지는 접근이 불가능합니다.",
				http.StatusForbidden,
			)
			return
		}

		req.Host = targetURL.Host
		proxy.ServeHTTP(writer, req)
	})

	log.Printf("http://localhost:8080 -> %s (via %s)", targetSite, proxyURL)
	log.Fatal(http.ListenAndServe(":8080", nil))
}