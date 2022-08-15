package samsung

import (
	"time"

	"github.com/kpeu3i/go-tizen-tv/tizenapi"
)

type UDPAPIClient interface {
	MAC() string
	Subnet() string
	Port() string
	WakeUp() error
}

type HTTPAPIClient interface {
	Host() string
	Port() string
	DialTimeout() time.Duration
	RequestTimeout() time.Duration
	ResponseTimeout() time.Duration
	IsAvailable() bool
	GetInfo() (tizenapi.GetInfoResponse, error)
	GetApp(id string) (tizenapi.GetAppResponse, error)
	OpenApp(id string) error
	InstallApp(id string) error
	CloseApp(id string) error
}

type WebsocketAPIClient interface {
	Host() string
	Port() string
	IsSecure() bool
	DialTimeout() time.Duration
	ReadTimeout() time.Duration
	WriteTimeout() time.Duration
	ClientID() string
	IsAvailable() bool
	IsConnected() bool
	Connect(token string) (tizenapi.ConnectResponseMessage, error)
	GetApps() (tizenapi.GetAppsResponseMessage, error)
	OpenApp(id string, actionType tizenapi.WebsocketOpenAppActionType, metaTag string) error
	SendKey(key string, state tizenapi.WebsocketKeyState) error
	Close() error
}
