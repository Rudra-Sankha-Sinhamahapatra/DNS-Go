package main

import (
	"fmt"
	"net"

	"github.com/google/gopacket/layers"
)

func serveDNS(u *net.UDPConn, clientAddr net.Addr, request *layers.DNS) {
	domain := string(request.Questions[0].Name)

	ip, found := records[domain]
	if !found {
		resolvedIP, err := resolveFromExternal(domain)
		if err != nil {
			fmt.Println("❌ Failed to resolve domain", domain)
			sendDNSResponse(u, clientAddr, request, "", true)
			return
		}

		records[domain] = resolvedIP
		ip = resolvedIP
		fmt.Println("✅ Domain exists on IP:", ip)
	} else {
		fmt.Println("✅ Domain found in cache:", ip)
	}
	sendDNSResponse(u, clientAddr, request, ip, false)
}
