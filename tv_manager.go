package samsung

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/kpeu3i/go-tizen-tv/ssdp"
	"github.com/kpeu3i/go-tizen-tv/tizenapi"
)

const (
	defaultClientID          = "GoTizenTV"
	defaultConfigStoragePath = "config.yaml"
)

type TVConfigStorage interface {
	Load() (TVManagerConfig, error)
	Store(config TVManagerConfig) error
}

type SSDPDiscoverer interface {
	Discover() ([]ssdp.Service, error)
}

type (
	SSDPDiscovererFactory     func(options ...ssdp.Option) SSDPDiscoverer
	UDPAPIClientFactory       func(mac string, options ...tizenapi.UDPAPIOption) UDPAPIClient
	HTTPAPIClientFactory      func(host string, options ...tizenapi.HTTPAPIOption) HTTPAPIClient
	WebsocketAPIClientFactory func(
		host string,
		clientID string,
		options ...tizenapi.WebsocketAPIOption,
	) WebsocketAPIClient
)

type TVManagerOption func(*TVManager)

func WithTVManagerConfigStorage(storage TVConfigStorage) TVManagerOption {
	return func(manager *TVManager) {
		manager.configStorage = storage
	}
}

func WithTVManagerSSDPDiscovererFactory(factory SSDPDiscovererFactory) TVManagerOption {
	return func(manager *TVManager) {
		manager.ssdpDiscovererFactory = factory
	}
}

func WithTVManagerUDPAPIClientFactory(factory UDPAPIClientFactory) TVManagerOption {
	return func(manager *TVManager) {
		manager.udpClientFactory = factory
	}
}

func WithTVManagerHTTPAPIClientFactory(factory HTTPAPIClientFactory) TVManagerOption {
	return func(manager *TVManager) {
		manager.httpClientFactory = factory
	}
}

func WithTVManagerWebsocketAPIClientFactory(factory WebsocketAPIClientFactory) TVManagerOption {
	return func(manager *TVManager) {
		manager.websocketClientFactory = factory
	}
}

type TVManager struct {
	configStorage          TVConfigStorage
	ssdpDiscovererFactory  SSDPDiscovererFactory
	udpClientFactory       UDPAPIClientFactory
	httpClientFactory      HTTPAPIClientFactory
	websocketClientFactory WebsocketAPIClientFactory
	ssdpDiscoverer         SSDPDiscoverer
}

func NewTVManager(options ...TVManagerOption) *TVManager {
	manager := &TVManager{
		configStorage: NewTVManagerConfigStorageYAML(defaultConfigStoragePath),
		ssdpDiscovererFactory: func(options ...ssdp.Option) SSDPDiscoverer {
			return ssdp.NewDiscoverer(options...)
		},
		udpClientFactory: func(mac string, options ...tizenapi.UDPAPIOption) UDPAPIClient {
			return tizenapi.NewUDPAPIClient(mac, options...)
		},
		httpClientFactory: func(host string, options ...tizenapi.HTTPAPIOption) HTTPAPIClient {
			return tizenapi.NewHTTPAPIClient(host, options...)
		},
		websocketClientFactory: func(
			host string,
			clientID string,
			options ...tizenapi.WebsocketAPIOption,
		) WebsocketAPIClient {
			return tizenapi.NewWebsocketAPIClient(host, clientID, options...)
		},
	}

	for _, option := range options {
		option(manager)
	}

	return manager
}

func (m *TVManager) Discover() ([]*TV, error) {
	hosts, err := m.discoverHosts()
	if err != nil {
		return nil, err
	}

	tvs := make([]*TV, 0, len(hosts))
	for _, host := range hosts {
		httpClient := m.httpClientFactory(host)
		if !httpClient.IsAvailable() {
			continue
		}

		info, err := httpClient.GetInfo()
		if err != nil {
			continue
		}

		device := TVDevice(info.Device)

		websocketClient := m.websocketClientFactory(
			device.IP(),
			defaultClientID,
			tizenapi.WithWebsocketIsSecure(true),
		)

		if !websocketClient.IsAvailable() {
			websocketClient = m.websocketClientFactory(device.IP(), defaultClientID)
			if !websocketClient.IsAvailable() {
				continue
			}
		}

		udpClient := m.udpClientFactory(device.MAC())

		deviceConfig := buildDeviceConfig(
			info.Device,
			udpClient,
			httpClient,
			websocketClient,
		)

		tvs = append(tvs, m.createTV(deviceConfig))
	}

	return tvs, nil
}

func (m *TVManager) Store(tvs ...*TV) error {
	if len(tvs) == 0 {
		return nil
	}

	config, err := m.loadConfig()
	if err != nil {
		return err
	}

	for _, tv := range tvs {
		info, err := tv.Info()
		if err != nil {
			return err
		}

		newDeviceConfig := buildDeviceConfig(
			info.Device,
			tv.udpClient,
			tv.httpClient,
			tv.websocketClient,
		)

		existingDeviceConfig, exists := config.DeviceConfig(info.Device.ID())
		if exists {
			newDeviceConfig.Name = existingDeviceConfig.Name
			newDeviceConfig.WebsocketAPI.ClientID = existingDeviceConfig.WebsocketAPI.ClientID
			newDeviceConfig.WebsocketAPI.Token = existingDeviceConfig.WebsocketAPI.Token
		}

		config.SetDeviceConfig(newDeviceConfig)
	}

	return m.storeConfig(config)
}

func (m *TVManager) Load() ([]*TV, error) {
	config, err := m.loadConfig()
	if err != nil {
		return nil, err
	}

	tvs := make([]*TV, 0, len(config.Devices))
	for _, deviceConfig := range config.Devices {
		tvs = append(tvs, m.createTV(deviceConfig))
	}

	return tvs, nil
}

func (m *TVManager) LoadByID(id string) (*TV, error) {
	config, err := m.loadConfig()
	if err != nil {
		return nil, err
	}

	for _, deviceConfig := range config.Devices {
		if deviceConfig.ID == id {
			return m.createTV(deviceConfig), nil
		}
	}

	return nil, fmt.Errorf("tv %s not found", id)
}

func (m *TVManager) discoverHosts() ([]string, error) {
	discoverer, err := m.discoverer()
	if err != nil {
		return nil, err
	}

	services, err := discoverer.Discover()
	if err != nil {
		return nil, err
	}

	var hosts []string

	exists := map[string]struct{}{}
	for _, service := range services {
		u, err := url.Parse(service.Location)
		if err != nil {
			continue
		}

		if _, ok := exists[u.Hostname()]; ok {
			continue
		}

		hosts = append(hosts, u.Hostname())
		exists[u.Hostname()] = struct{}{}
	}

	return hosts, nil
}

func (m *TVManager) discoverer() (SSDPDiscoverer, error) {
	if m.ssdpDiscoverer == nil {
		config, err := m.loadConfig()
		if err != nil {
			return nil, err
		}

		m.ssdpDiscoverer = m.ssdpDiscovererFactory(
			ssdp.WithSearchDuration(config.Discovery.Duration),
		)
	}

	return m.ssdpDiscoverer, nil
}

func (m *TVManager) createTV(deviceConfig DeviceConfig) *TV {
	udpClientOptions := []tizenapi.UDPAPIOption{
		tizenapi.WithUDPSubnet(deviceConfig.UDPAPI.Subnet),
		tizenapi.WithUDPPort(deviceConfig.UDPAPI.Port),
	}

	httpClientOptions := []tizenapi.HTTPAPIOption{
		tizenapi.WithHTTPPort(deviceConfig.HTTPAPI.Port),
		tizenapi.WithHTTPDialTimeout(deviceConfig.HTTPAPI.DialTimeout),
		tizenapi.WithHTTPRequestHeaderTimeout(deviceConfig.HTTPAPI.RequestTimeout),
		tizenapi.WithHTTPResponseTimeout(deviceConfig.HTTPAPI.ResponseTimeout),
	}

	websocketClientOptions := []tizenapi.WebsocketAPIOption{
		tizenapi.WithWebsocketIsSecure(deviceConfig.WebsocketAPI.IsSecure),
		tizenapi.WithWebsocketPort(deviceConfig.WebsocketAPI.Port),
		tizenapi.WithWebsocketReadTimeout(deviceConfig.WebsocketAPI.ReadTimeout),
		tizenapi.WithWebsocketWriteTimeout(deviceConfig.WebsocketAPI.WriteTimeout),
	}

	tv := NewTV(
		m.udpClientFactory(deviceConfig.MAC, udpClientOptions...),
		m.httpClientFactory(deviceConfig.Host, httpClientOptions...),
		m.websocketClientFactory(deviceConfig.Host, deviceConfig.WebsocketAPI.ClientID, websocketClientOptions...),
		deviceConfig.WebsocketAPI.Token,
	)

	tv.OnAuthorize(func(token string) error {
		config, err := m.loadConfig()
		if err != nil {
			return err
		}

		deviceConfig, exists := config.DeviceConfig(deviceConfig.ID)
		if !exists {
			return errors.New(fmt.Sprintf("configuration for device %s is not provided", deviceConfig.ID))
		}

		deviceConfig.WebsocketAPI.Token = token

		config.SetDeviceConfig(deviceConfig)

		return m.storeConfig(config)
	})

	return tv
}

func (m *TVManager) loadConfig() (TVManagerConfig, error) {
	config, err := m.configStorage.Load()
	if err != nil {
		return TVManagerConfig{}, err
	}

	if config.IsEmpty() {
		config.Discovery.Duration = 5 * time.Second
	}

	return config, nil
}

func (m *TVManager) storeConfig(config TVManagerConfig) error {
	return m.configStorage.Store(config)
}

func buildDeviceConfig(
	device TVDevice,
	udpClient UDPAPIClient,
	httpClient HTTPAPIClient,
	websocketClient WebsocketAPIClient,
) DeviceConfig {
	deviceConfig := DeviceConfig{
		ID:   device.ID(),
		Name: device.Name(),
		Host: device.IP(),
		MAC:  device.MAC(),
	}

	deviceConfig.UDPAPI.Subnet = udpClient.Subnet()
	deviceConfig.UDPAPI.Port = udpClient.Port()

	deviceConfig.HTTPAPI.Port = httpClient.Port()
	deviceConfig.HTTPAPI.DialTimeout = httpClient.DialTimeout()
	deviceConfig.HTTPAPI.RequestTimeout = httpClient.RequestTimeout()
	deviceConfig.HTTPAPI.ResponseTimeout = httpClient.ResponseTimeout()

	deviceConfig.WebsocketAPI.Port = websocketClient.Port()
	deviceConfig.WebsocketAPI.IsSecure = websocketClient.IsSecure()
	deviceConfig.WebsocketAPI.DialTimeout = websocketClient.DialTimeout()
	deviceConfig.WebsocketAPI.ReadTimeout = websocketClient.ReadTimeout()
	deviceConfig.WebsocketAPI.WriteTimeout = websocketClient.WriteTimeout()
	deviceConfig.WebsocketAPI.ClientID = websocketClient.ClientID()

	return deviceConfig
}
