package ws

import (
	"net"
	"net/url"
)

func parseHost(rawurl string) (host string) {
	defer func() {
		domain, _, err := net.SplitHostPort(host)
		if err == nil {
			host = domain
		}
	}()
	u, err := url.ParseRequestURI(rawurl)
	if err != nil || u.Host == "" {
		u, err = url.ParseRequestURI("https://" + rawurl)
		if err != nil {
			host = rawurl
			return
		}
	}
	host = u.Host
	return
}

// areUrlEqual returns true if the provided URLs describe the same endpoint.
func areUrlEqual(urlStr1, urlStr2 string) bool {

	// Resolve hosts.
	host1 := parseHost(urlStr1)
	host2 := parseHost(urlStr2)

	// Try to find a matching address when looking hosts up.
	addrs := make(map[string]int)
	addrs1, err := net.LookupHost(host1)
	if err == nil {
		for _, addr := range addrs1 {
			addrs[addr]++
		}
	}
	addrs2, err := net.LookupHost(host2)
	if err == nil {
		for _, addr := range addrs2 {
			addrs[addr]++
			if addrs[addr] > 1 {
				return true
			}
		}
	}

	// No match return false.
	return false
}
