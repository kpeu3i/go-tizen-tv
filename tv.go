package samsung

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/kpeu3i/go-tizen-tv/tizenapi"
)

const (
	defaultAppBrowser      = "org.tizen.browser"
	defaultTimeoutPowerOn  = 15 * time.Second
	defaultTimeoutPowerOff = 1 * time.Minute
)

type AuthorizeHandler func(token string) error

type TVOption func(*TV)

func WithPowerOnTimeout(timeout time.Duration) TVOption {
	return func(tv *TV) {
		tv.powerOnTimeout = timeout
	}
}

func WithPowerOffTimeout(timeout time.Duration) TVOption {
	return func(tv *TV) {
		tv.powerOffTimeout = timeout
	}
}

func WithKeyPowerKey(key Key) TVOption {
	return func(tv *TV) {
		tv.keyPowerOff = key
	}
}

type TVApp struct {
	ID        string
	Name      string
	IsRunning bool
	IsVisible bool
	Version   string
}

type TVInfo struct {
	ID        string
	Type      string
	Name      string
	Version   string
	URI       string
	Remote    string
	Device    TVDevice
	IsSupport map[string]string
}

type TV struct {
	udpClient        UDPAPIClient
	httpClient       HTTPAPIClient
	websocketClient  WebsocketAPIClient
	token            string
	keyPowerOff      Key
	powerOnTimeout   time.Duration
	powerOffTimeout  time.Duration
	authorizeHandler AuthorizeHandler
}

func NewTV(
	udpAPIClient UDPAPIClient,
	httpAPIClient HTTPAPIClient,
	websocketAPIClient WebsocketAPIClient,
	token string,
	options ...TVOption,
) *TV {
	tv := &TV{
		udpClient:       udpAPIClient,
		httpClient:      httpAPIClient,
		websocketClient: websocketAPIClient,
		token:           token,
		keyPowerOff:     KEY_POWER,
		powerOnTimeout:  defaultTimeoutPowerOn,
		powerOffTimeout: defaultTimeoutPowerOff,
	}

	for _, option := range options {
		option(tv)
	}

	return tv
}

func (tv *TV) PowerOn() error {
	ctx, cancel := context.WithTimeout(context.Background(), tv.powerOnTimeout)
	defer cancel()

	ready := make(chan struct{}, 1)
	go func() {
		defer func() {
			ready <- struct{}{}
		}()

		for {
			if tv.IsReady() {
				return
			}

			err := tv.udpClient.WakeUp()
			if err != nil {
				continue
			}

			if ctx.Err() != nil {
				return
			}

			time.Sleep(3 * time.Second)
		}
	}()

	select {
	case <-ready:
		err := tv.establishWebsocketConnection()
		if err != nil {
			return err
		}
	case <-ctx.Done():
		return errors.New("cannot find TV in the network")
	}

	return nil
}

func (tv *TV) PowerOff() error {
	ctx, cancel := context.WithTimeout(context.Background(), tv.powerOffTimeout)
	defer cancel()

	if !tv.isAvailable() {
		return nil
	}

	err := tv.ClickKey(tv.keyPowerOff)
	if err != nil {
		return err
	}

	ready := make(chan struct{}, 1)
	go func() {
		defer func() {
			ready <- struct{}{}
		}()

		for {
			if !tv.isAvailable() {
				break
			}

			time.Sleep(1 * time.Second)
		}
	}()

	select {
	case <-ready:
		return nil
	case <-ctx.Done():
		return errors.New("unable to power off a TV")
	}
}

func (tv *TV) OnAuthorize(handler AuthorizeHandler) {
	tv.authorizeHandler = handler
}

func (tv *TV) IsReady() bool {
	if !tv.isAvailable() {
		return false
	}

	_, err := tv.httpClient.GetInfo()
	if err != nil {
		return false
	}

	return true
}

func (tv *TV) Info() (TVInfo, error) {
	response, err := tv.httpClient.GetInfo()
	if err != nil {
		return TVInfo{}, err
	}

	var isSupport map[string]string
	err = json.Unmarshal([]byte(response.IsSupport), &isSupport)
	if err != nil {
		return TVInfo{}, err
	}

	info := TVInfo{
		ID:        response.ID,
		Type:      response.Type,
		Name:      response.Name,
		Version:   response.Version,
		URI:       response.URI,
		Remote:    response.Remote,
		Device:    response.Device,
		IsSupport: isSupport,
	}

	return info, nil
}

func (tv *TV) Apps() ([]TVApp, error) {
	err := tv.ensureWebsocketConnection()
	if err != nil {
		return nil, err
	}

	appsResponse, err := tv.websocketClient.GetApps()
	if err != nil {
		return nil, err
	}

	apps := make([]TVApp, 0, len(appsResponse.Data.Data))
	for _, item := range appsResponse.Data.Data {
		appResponse, err := tv.httpClient.GetApp(item.AppId)
		if err != nil {
			continue
		}

		app := TVApp{
			ID:        item.AppId,
			Name:      item.Name,
			IsRunning: appResponse.Running,
			IsVisible: appResponse.Visible,
			Version:   appResponse.Version,
		}

		apps = append(apps, app)
	}

	return apps, nil
}

func (tv *TV) App(id string) (TVApp, error) {
	response, err := tv.httpClient.GetApp(id)
	if err != nil {
		return TVApp{}, err
	}

	app := TVApp{
		ID:        response.ID,
		Name:      response.Name,
		IsRunning: response.Running,
		IsVisible: response.Visible,
		Version:   response.Version,
	}

	return app, nil
}

func (tv *TV) OpenApp(id string) error {
	return tv.httpClient.OpenApp(id)
}

func (tv *TV) InstallApp(id string) error {
	return tv.httpClient.InstallApp(id)
}

func (tv *TV) CloseApp(id string) error {
	return tv.httpClient.CloseApp(id)
}

func (tv *TV) OpenBrowser(url string) error {
	err := tv.ensureWebsocketConnection()
	if err != nil {
		return err
	}

	err = tv.websocketClient.OpenApp(
		defaultAppBrowser,
		tizenapi.WebsocketOpenAppActionTypeNativeLaunch,
		url,
	)
	if err != nil {
		return err
	}

	return nil
}

func (tv *TV) ClickKey(key Key) error {
	err := tv.ensureWebsocketConnection()
	if err != nil {
		return err
	}

	return tv.websocketClient.SendKey(string(key), tizenapi.WebsocketKeyStateClick)
}

func (tv *TV) PressKey(key Key) error {
	err := tv.ensureWebsocketConnection()
	if err != nil {
		return err
	}

	return tv.websocketClient.SendKey(string(key), tizenapi.WebsocketKeyStatePress)
}

func (tv *TV) ReleaseKey(key Key) error {
	err := tv.ensureWebsocketConnection()
	if err != nil {
		return err
	}

	return tv.websocketClient.SendKey(string(key), tizenapi.WebsocketKeyStateRelease)
}

func (tv *TV) SendKeys(sequence KeySequence) error {
	err := tv.ensureWebsocketConnection()
	if err != nil {
		return err
	}

	for _, command := range sequence {
		switch command.action {
		case keyActionClick:
			err := tv.ClickKey(command.key)
			if err != nil {
				return err
			}
		case keyActionPress:
			err := tv.PressKey(command.key)
			if err != nil {
				return err
			}
		case keyActionRelease:
			err := tv.ReleaseKey(command.key)
			if err != nil {
				return err
			}
		}

		if command.wait > 0 {
			time.Sleep(command.wait)
		}
	}

	return nil
}

func (tv *TV) Close() error {
	return tv.websocketClient.Close()
}

func (tv *TV) isAvailable() bool {
	return tv.httpClient.IsAvailable() && tv.websocketClient.IsAvailable()
}

func (tv *TV) ensureWebsocketConnection() error {
	if !tv.websocketClient.IsConnected() {
		err := tv.establishWebsocketConnection()
		if err != nil {
			return err
		}
	}

	return nil
}

func (tv *TV) establishWebsocketConnection() error {
	response, err := tv.websocketClient.Connect(tv.token)
	if err != nil {
		return err
	}

	if response.Data.Token != "" {
		tv.token = response.Data.Token
		if tv.authorizeHandler != nil {
			err := tv.authorizeHandler(response.Data.Token)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
