package main

import (
	"example/user/hello/src/utils"
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func main() {
	utils.LoadConfig()
	utils.InitLogger()
	Ip := utils.AppConfig.Ip
	utils.Logger.Infof("DNS Server Started on %s:%d", Ip, utils.AppConfig.ServerPort)
	fmt.Printf("DNS Server Started on %s:%d\n", Ip, utils.AppConfig.ServerPort)
	records = make(map[string]string)

	addr := net.UDPAddr{
		Port: utils.AppConfig.ServerPort,
		IP:   net.ParseIP(utils.AppConfig.Ip),
	}

	u, err := net.ListenUDP("udp", &addr)
	if err != nil {
		utils.Logger.Error("❌ Error setting up UDP server:", err)
		return
	}
	defer u.Close()

	// Wait to get requests on that port
	for {
		tmp := make([]byte, 1024)
		_, clientAddr, err := u.ReadFrom(tmp)
		if err != nil {
			utils.Logger.Error("❌ Error reading UDP packet:", err)
			continue
		}

		// Parse the packet as a DNS layer
		packet := gopacket.NewPacket(tmp, layers.LayerTypeDNS, gopacket.Default)
		dnsPacket := packet.Layer(layers.LayerTypeDNS)
		if dnsPacket == nil {
			utils.Logger.Warn("Non-DNS packet received")
			continue
		}

		tcp, _ := dnsPacket.(*layers.DNS)
		if tcp == nil {
			utils.Logger.Warn("Failed to parse DNS packet")
			continue
		}

		// Call the function to serve the DNS query
		serveDNS(u, clientAddr, tcp)
	}
}
