package logger

import (
	"fmt"
	"net"
)

type UdpSyncer struct {
	conn *net.UDPConn
}

func NewUDPSyncer(bindIp string, bindPort int) *UdpSyncer {
	// ResolveUDPAddr returns an address of UDP end point.
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", bindIp, bindPort))
	if err != nil {
		fmt.Println("Failed to resolve address", err)
	}

	// DialUDP connects to the remote address raddr on the network net
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Failed to dial address", err)
	}

	return &UdpSyncer{conn: conn}
}

func (s *UdpSyncer) Write(p []byte) (n int, err error) {
	return s.conn.Write(p)
}

func (s *UdpSyncer) Sync() error {
	return nil
}


