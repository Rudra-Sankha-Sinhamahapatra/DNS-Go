package main

import "net"

func resolveFromExternal(domain string) (string, error) {
	ips, err := net.LookupIP(domain)
	if err != nil || len(ips) == 0 {
		return "", err
	}
	return ips[0].String(), nil
}
