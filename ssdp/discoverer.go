package ssdp

import (
	"time"

	"github.com/koron/go-ssdp"
)

type SearchType string

const (
	SearchAll        SearchType = "ssdp:all"        // Search all services and devices
	SearchRootDevice SearchType = "upnp:rootdevice" // Search UPnP root devices

	defaultSearchType     = SearchRootDevice
	defaultSearchDuration = 5 * time.Second
)

type Option func(discoverer *Discoverer)

func WithSearchDuration(duration time.Duration) Option {
	return func(d *Discoverer) {
		d.searchDuration = duration
	}
}

func WithSearchType(searchType SearchType) Option {
	return func(d *Discoverer) {
		d.searchType = searchType
	}
}

type Service struct {
	Type     string
	USN      string
	Location string
	Server   string
}

type Discoverer struct {
	searchDuration time.Duration
	searchType     SearchType
}

func NewDiscoverer(options ...Option) *Discoverer {
	discoverer := &Discoverer{
		searchDuration: defaultSearchDuration,
		searchType:     defaultSearchType,
	}

	for _, option := range options {
		option(discoverer)
	}

	return discoverer
}

func (d *Discoverer) Discover() ([]Service, error) {
	sec := int(d.searchDuration / time.Second)

	response, err := ssdp.Search(string(d.searchType), sec, "")
	if err != nil {
		return nil, err
	}

	services := make([]Service, 0, len(response))
	for _, srv := range response {
		service := Service{
			Type:     srv.Type,
			USN:      srv.USN,
			Location: srv.Location,
			Server:   srv.Server,
		}

		services = append(services, service)
	}

	return services, nil
}
