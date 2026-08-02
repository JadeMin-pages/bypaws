package web

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Run() {
	target, _ := url.Parse(targetSite)

	proxyURL, err := GetFirstProxy()
	if err != nil {
		log.Fatalf("프록시 로드 실패: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.Transport = &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		req.Host = target.Host
		proxy.ServeHTTP(writer, req)
	})

	log.Printf("http://localhost:8080 -> %s (via %s)", targetSite, proxyURL)
	log.Fatal(http.ListenAndServe(":8080", nil))
}