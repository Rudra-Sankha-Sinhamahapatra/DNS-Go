package main

import (
	"example/user/hello/src/utils"
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
			utils.Logger.Error("❌ Failed to resolve domain", domain)
			fmt.Println("❌ Failed to resolve domain", domain)
			sendDNSResponse(u, clientAddr, request, "", true)
			return
		}

		records[domain] = resolvedIP
		ip = resolvedIP
		utils.Logger.Info("✅ Domain ", domain, " exists on IP:", ip)
		fmt.Println("✅ Domain ", domain, " exists on IP:", ip)
	} else {
		utils.Logger.Info("✅ Domain ", domain, " found in cache:", ip)
		fmt.Println("✅ Domain ", domain, " found in cache:", ip)
	}
	sendDNSResponse(u, clientAddr, request, ip, false)
}
