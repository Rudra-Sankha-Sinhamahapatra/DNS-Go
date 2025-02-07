package main

import (
	"example/user/hello/src/utils"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func sendDNSResponse(u *net.UDPConn, clientAddr net.Addr, request *layers.DNS, ip string, notFound bool) {
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

	if notFound || ip == "" {
		replyMess.ResponseCode = layers.DNSResponseCodeNXDomain
	} else {

		dnsAnswer := layers.DNSResourceRecord{
			Name:  request.Questions[0].Name,
			Type:  layers.DNSTypeA,
			Class: layers.DNSClassIN,
			IP:    net.ParseIP(ip),
			TTL:   300,
		}
		replyMess.ANCount = 1
		replyMess.Answers = append(replyMess.Answers, dnsAnswer)
	}

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{}
	err := replyMess.SerializeTo(buf, opts)
	if err != nil {
		utils.Logger.Error("Error serializing Dns Response ", err)
		return
	}

	_, err = u.WriteTo(buf.Bytes(), clientAddr)
	if err != nil {
		utils.Logger.Error("Error Seding DNS response: ", err)
	}
}
