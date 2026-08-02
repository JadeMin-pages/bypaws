package proxy

import (
	"fmt"
	"net/http"
	"net/url"
	"encoding/json"
)

type Proxy struct {
	IP			string	`json:"ip"`
	Port		int		`json:"port"`
	Protocol	string	`json:"protocol"`
}

const proxyListURL = "https://raw.githubusercontent.com/proxifly/free-proxy-list/refs/heads/main/proxies/countries/JP/data.json"

// ponytail: 검증 없이 목록의 첫 번째 프록시를 즉시 반환. 해당 IP가 죽어있으면 접속 에러 발생 (업그레이드: 연결 테스트/순회 추가).
func GetFirstProxy() (*url.URL, error) {
	resp, err := http.Get(proxyListURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var list []Proxy
	err = json.NewDecoder(resp.Body).Decode(&list)
	if err != nil || len(list) == 0 {
		return nil, fmt.Errorf("프록시 목록이 비어있거나 파싱 실패")
	}

	return url.Parse(
		fmt.Sprintf(
			"%s://%s:%d",
			list[0].Protocol, list[0].IP, list[0].Port,
		),
	)
}