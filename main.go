package main

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var records map[string]string

func main() {
	fmt.Println("DNS Server Started")
	records = map[string]string{
		"google.com": "216.58.196.142",
		"amazon.com": "176.32.103.205",
	}

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

func serveDNS(u *net.UDPConn, clientAddr net.Addr, request *layers.DNS) {
	replyMess := layers.DNS{
		ID:           request.ID,            // Copy query ID from request
		QR:           true,                  // Indicate response
		OpCode:       layers.DNSOpCodeQuery, // Query response
		AA:           true,                  // Authoritative Answer
		TC:           false,                 // Truncated
		RD:           request.RD,            // Recursion Desired
		RA:           true,                  // Recursion Available
		ResponseCode: layers.DNSResponseCodeNoErr,
	}

	// Process DNS Questions and construct DNS Answers
	var dnsAnswer layers.DNSResourceRecord
	dnsAnswer.Type = layers.DNSTypeA
	var ip string
	var ok bool

	// Look up the query domain name in the records map
	ip, ok = records[string(request.Questions[0].Name)]
	if !ok {
		// Domain not found, return Name Error (NXDomain)
		replyMess.ResponseCode = layers.DNSResponseCodeNXDomain
	} else {
		// Successfully found IP, convert string to net.IP
		dnsAnswer.IP = net.ParseIP(ip)
		dnsAnswer.Name = request.Questions[0].Name
		dnsAnswer.Class = layers.DNSClassIN

		// Set Answer Count
		replyMess.ANCount = 1
		// Add the answer record
		replyMess.Answers = append(replyMess.Answers, dnsAnswer)
	}

	// Serialize the response
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{}
	err := replyMess.SerializeTo(buf, opts)
	if err != nil {
		fmt.Println("Error serializing DNS response:", err)
		return
	}

	// Send the response back to the client
	_, err = u.WriteTo(buf.Bytes(), clientAddr)
	if err != nil {
		fmt.Println("Error sending DNS response:", err)
	}
}
