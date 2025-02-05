package main

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func main() {
	fmt.Println("DNS Server Started on 127.0.0.1:8090")
	records = make(map[string]string)

	// Listen on UDP Port 8090
	addr := net.UDPAddr{
		Port: 8090,
		IP:   net.ParseIP("127.0.0.1"),
	}

	u, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error setting up UDP server:", err)
		return
	}
	defer u.Close()

	// Wait to get requests on that port
	for {
		tmp := make([]byte, 1024)
		_, clientAddr, err := u.ReadFrom(tmp)
		if err != nil {
			fmt.Println("Error reading UDP packet:", err)
			continue
		}

		// Parse the packet as a DNS layer
		packet := gopacket.NewPacket(tmp, layers.LayerTypeDNS, gopacket.Default)
		dnsPacket := packet.Layer(layers.LayerTypeDNS)
		if dnsPacket == nil {
			fmt.Println("Non-DNS packet received")
			continue
		}

		tcp, _ := dnsPacket.(*layers.DNS)
		if tcp == nil {
			fmt.Println("Failed to parse DNS packet")
			continue
		}

		// Call the function to serve the DNS query
		serveDNS(u, clientAddr, tcp)
	}
}
