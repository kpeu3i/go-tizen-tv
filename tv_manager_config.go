package samsung

import (
	"time"
)

type DeviceConfig struct {
	ID     string `json:"id" yaml:"id"`
	Name   string `json:"name" yaml:"name"`
	Host   string `json:"host" yaml:"host"`
	MAC    string `json:"mac" yaml:"mac"`
	UDPAPI struct {
		Subnet string `json:"subnet" yaml:"subnet"`
		Port   string `json:"port" yaml:"port"`
	} `json:"udp_api" yaml:"udp_api"`
	HTTPAPI struct {
		Port            string        `json:"port" yaml:"port"`
		DialTimeout     time.Duration `json:"dial_timeout" yaml:"dial_timeout"`
		RequestTimeout  time.Duration `json:"request_timeout" yaml:"request_timeout"`
		ResponseTimeout time.Duration `json:"response_timeout" yaml:"response_timeout"`
	} `json:"http_api" yaml:"http_api"`
	WebsocketAPI struct {
		Port         string        `json:"port" yaml:"port"`
		IsSecure     bool          `json:"is_secure" yaml:"is_secure"`
		DialTimeout  time.Duration `json:"dial_timeout" yaml:"dial_timeout"`
		ReadTimeout  time.Duration `json:"read_timeout" yaml:"read_timeout"`
		WriteTimeout time.Duration `json:"write_timeout" yaml:"write_timeout"`
		ClientID     string        `json:"client_id" yaml:"client_id"`
		Token        string        `json:"token" yaml:"token"`
	} `json:"websocket_api" yaml:"websocket_api"`
}

type TVManagerConfig struct {
	Discovery struct {
		Duration time.Duration `json:"duration" yaml:"duration"`
	} `json:"discovery" yaml:"discovery"`
	Devices []DeviceConfig `json:"devices" yaml:"devices"`
}

func (c *TVManagerConfig) IsEmpty() bool {
	return c.Discovery.Duration == 0 && c.Devices == nil
}

func (c *TVManagerConfig) DeviceConfig(id string) (DeviceConfig, bool) {
	for _, device := range c.Devices {
		if device.ID == id {
			return device, true
		}
	}

	return DeviceConfig{}, false
}

func (c *TVManagerConfig) SetDeviceConfig(deviceConfig DeviceConfig) {
	for i, device := range c.Devices {
		if device.ID == deviceConfig.ID {
			c.Devices[i] = deviceConfig

			return
		}
	}

	c.Devices = append(c.Devices, deviceConfig)
}
