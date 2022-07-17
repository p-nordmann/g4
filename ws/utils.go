/*
G4 is an open-source board game inspired by the popular game of connect-4.
Copyright (C) 2022  Pierre-Louis Nordmann

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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
