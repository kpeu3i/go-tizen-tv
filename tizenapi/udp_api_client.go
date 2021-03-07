package tizenapi

import (
	"errors"
	"fmt"
	"net"
)

const (
	defaultUDPSubnet = "255.255.255.255"
	defaultUDPPort   = "9"
)

type wolPacket [102]byte

type UDPAPIOption func(*UDPAPIClient)

func WithUDPSubnet(subnet string) UDPAPIOption {
	return func(client *UDPAPIClient) {
		client.subnet = subnet
	}
}

func WithUDPPort(port string) UDPAPIOption {
	return func(client *UDPAPIClient) {
		client.port = port
	}
}

type UDPAPIClient struct {
	mac    string
	subnet string
	port   string
}

func NewUDPAPIClient(mac string, options ...UDPAPIOption) *UDPAPIClient {
	client := &UDPAPIClient{
		mac:    mac,
		subnet: defaultUDPSubnet,
		port:   defaultUDPPort,
	}

	for _, option := range options {
		option(client)
	}

	return client
}

func (s *UDPAPIClient) MAC() string {
	return s.mac
}

func (s *UDPAPIClient) Subnet() string {
	return s.subnet
}

func (s *UDPAPIClient) Port() string {
	return s.port
}

func (s *UDPAPIClient) WakeUp() error {
	mac, err := s.parseMAC()
	if err != nil {
		return err
	}

	packet := constructWOLPacket(mac)

	err = s.broadcastWOLPacket(packet)
	if err != nil {
		return err
	}

	return nil
}

func (s *UDPAPIClient) broadcastWOLPacket(packet wolPacket) error {
	connection, err := net.Dial("udp", fmt.Sprintf("%s:%s", s.subnet, s.port))
	if err != nil {
		return err
	}

	defer func() {
		_ = connection.Close()
	}()

	_, err = connection.Write(packet[:])
	if err != nil {
		return err
	}

	return nil
}

func (s *UDPAPIClient) parseMAC() (net.HardwareAddr, error) {
	mac, err := net.ParseMAC(s.mac)
	if err != nil {
		return nil, fmt.Errorf("parse MAC error: %s", err.Error())
	}

	if len(mac) != 6 {
		return nil, errors.New("parse MAC error: invalid EUI-48 MAC address")
	}

	return mac, nil
}

func constructWOLPacket(mac net.HardwareAddr) wolPacket {
	var packet wolPacket

	offset := 6
	copy(packet[0:], []byte{255, 255, 255, 255, 255, 255})

	for i := 0; i < 16; i++ {
		copy(packet[offset:], mac)
		offset += 6
	}

	return packet
}
