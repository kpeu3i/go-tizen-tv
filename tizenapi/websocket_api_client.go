package tizenapi

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"sort"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

const (
	defaultWebsocketPort         = "8002"
	defaultWebsocketSecurePort   = "8002"
	defaultWebsocketDialTimeout  = 1 * time.Second
	defaultWebsocketReadTimeout  = 5 * time.Second
	defaultWebsocketWriteTimeout = 5 * time.Second
)

type WebsocketKeyState string

const (
	WebsocketKeyStateClick   WebsocketKeyState = "Click"
	WebsocketKeyStatePress   WebsocketKeyState = "Press"
	WebsocketKeyStateRelease WebsocketKeyState = "Release"
)

type WebsocketOpenAppActionType string

const (
	WebsocketOpenAppActionTypeDeepLink     WebsocketOpenAppActionType = "DEEP_LINK"
	WebsocketOpenAppActionTypeNativeLaunch WebsocketOpenAppActionType = "NATIVE_LAUNCH"
)

type WebsocketAPIOption struct {
	setter   func(*WebsocketAPIClient)
	priority int
}

func WithWebsocketPort(port string) WebsocketAPIOption {
	return WebsocketAPIOption{
		setter: func(client *WebsocketAPIClient) {
			client.port = port
		},
		priority: 2,
	}
}

func WithWebsocketIsSecure(isSecure bool) WebsocketAPIOption {
	return WebsocketAPIOption{
		setter: func(client *WebsocketAPIClient) {
			client.isSecure = isSecure
			if isSecure {
				client.port = defaultWebsocketSecurePort
			} else {
				client.port = defaultWebsocketPort
			}
		},
		priority: 1,
	}
}

func WithWebsocketReadTimeout(timeout time.Duration) WebsocketAPIOption {
	return WebsocketAPIOption{
		setter: func(client *WebsocketAPIClient) {
			client.readTimeout = timeout
		},
		priority: 2,
	}
}

func WithWebsocketWriteTimeout(timeout time.Duration) WebsocketAPIOption {
	return WebsocketAPIOption{
		setter: func(client *WebsocketAPIClient) {
			client.writeTimeout = timeout
		},
		priority: 2,
	}
}

type WebsocketAPIClient struct {
	host             string
	port             string
	isSecure         bool
	dialTimeout      time.Duration
	readTimeout      time.Duration
	writeTimeout     time.Duration
	clientID         string
	connection       *websocket.Conn
	readerState      int32
	writerState      int32
	requestMessages  chan []byte
	responseMessages chan []byte
	quit             chan struct{}
}

func NewWebsocketAPIClient(
	host string,
	clientID string,
	options ...WebsocketAPIOption,
) *WebsocketAPIClient {
	client := &WebsocketAPIClient{
		host:             host,
		port:             defaultWebsocketPort,
		dialTimeout:      defaultWebsocketDialTimeout,
		readTimeout:      defaultWebsocketReadTimeout,
		writeTimeout:     defaultWebsocketWriteTimeout,
		clientID:         clientID,
		requestMessages:  make(chan []byte),
		responseMessages: make(chan []byte),
	}

	sort.Slice(options, func(i, j int) bool {
		return options[i].priority < options[j].priority
	})

	for _, option := range options {
		option.setter(client)
	}

	return client
}

func (c *WebsocketAPIClient) Host() string {
	return c.host
}

func (c *WebsocketAPIClient) Port() string {
	return c.port
}

func (c *WebsocketAPIClient) IsSecure() bool {
	return c.isSecure
}

func (c *WebsocketAPIClient) DialTimeout() time.Duration {
	return c.dialTimeout
}

func (c *WebsocketAPIClient) ReadTimeout() time.Duration {
	return c.readTimeout
}

func (c *WebsocketAPIClient) WriteTimeout() time.Duration {
	return c.writeTimeout
}

func (c *WebsocketAPIClient) ClientID() string {
	return c.clientID
}

func (c *WebsocketAPIClient) IsAvailable() bool {
	connection, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", c.host, c.port), c.dialTimeout)
	if err != nil {
		return false
	}

	defer func() {
		_ = connection.Close()
	}()

	return true
}

func (c *WebsocketAPIClient) IsConnected() bool {
	return atomic.LoadInt32(&c.readerState) == 1 || atomic.LoadInt32(&c.writerState) == 1
}

func (c *WebsocketAPIClient) Connect(token string) (ConnectResponseMessage, error) {
	if c.IsConnected() {
		return ConnectResponseMessage{}, errors.New("connection has been already opened")
	}

	dialer := *websocket.DefaultDialer
	if c.isSecure {
		dialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	dialer.NetDial = func(network, addr string) (net.Conn, error) {
		return net.DialTimeout(network, addr, c.dialTimeout)
	}

	connection, _, err := dialer.Dial(c.generateURL(token), nil)
	if err != nil {
		return ConnectResponseMessage{}, err
	}

	c.connection = connection
	c.quit = make(chan struct{})

	go func() {
		defer func() {
			_ = c.connection.Close()
		}()

		_ = c.runReader()
	}()

	go func() {
		defer func() {
			_ = c.connection.Close()
		}()

		_ = c.runWriter()
	}()

	message, err := c.wait()
	if err != nil {
		return ConnectResponseMessage{}, err
	}

	response := ConnectResponseMessage{}

	err = json.Unmarshal(message, &response)
	if err != nil {
		return ConnectResponseMessage{}, err
	}

	return response, nil
}

func (c *WebsocketAPIClient) GetApps() (GetAppsResponseMessage, error) {
	request := &GetAppsRequestMessage{}
	request.Method = "ms.channel.emit"
	request.Params.Event = "ed.installedApp.get"
	request.Params.To = "host"

	requestMessage, err := json.Marshal(request)
	if err != nil {
		return GetAppsResponseMessage{}, err
	}

	responseMessage, err := c.sendAndWait(requestMessage)
	if err != nil {
		return GetAppsResponseMessage{}, err
	}

	response := GetAppsResponseMessage{}
	err = json.Unmarshal(responseMessage, &response)
	if err != nil {
		return GetAppsResponseMessage{}, err
	}

	return response, nil
}

func (c *WebsocketAPIClient) OpenApp(id string, actionType WebsocketOpenAppActionType, metaTag string) error {
	request := &OpenAppRequestMessage{}
	request.Method = "ms.channel.emit"
	request.Params.Event = "ed.apps.launch"
	request.Params.To = "host"
	request.Params.Data.ActionType = string(actionType)
	request.Params.Data.AppID = id
	request.Params.Data.MetaTag = metaTag

	requestMessage, err := json.Marshal(request)
	if err != nil {
		return err
	}

	responseMessage, err := c.sendAndWait(requestMessage)
	if err != nil {
		return err
	}

	response := map[string]interface{}{}
	err = json.Unmarshal(responseMessage, &response)
	if err != nil {
		return err
	}

	isResponseValid := false
	if v, ok := response["data"].(bool); ok && v {
		isResponseValid = true
	}

	if _, ok := response["event"].(string); ok {
		isResponseValid = true
	}

	if !isResponseValid {
		return fmt.Errorf("invalid response: %s", string(responseMessage))
	}

	return nil
}

func (c *WebsocketAPIClient) SendKey(key string, state WebsocketKeyState) error {
	request := &SendKeyRequestMessage{}
	request.Method = "ms.remote.control"
	request.Params.Cmd = string(state)
	request.Params.DataOfCmd = key
	request.Params.Option = "false"
	request.Params.TypeOfRemote = "SendRemoteKey"

	requestData, err := json.Marshal(request)
	if err != nil {
		return err
	}

	err = c.send(requestData)
	if err != nil {
		return err
	}

	return nil
}

func (c *WebsocketAPIClient) Close() error {
	if c.connection == nil {
		return nil
	}

	close(c.quit)
	_ = c.connection.Close()

	maxAttempts := 10
	delay := 1 * time.Second

	attempts := 0
	for {
		attempts++

		if !c.IsConnected() {
			return nil
		}

		time.Sleep(delay)

		if attempts == maxAttempts {
			return errors.New("unable to close connection")
		}
	}
}

func (c *WebsocketAPIClient) runReader() error {
	atomic.StoreInt32(&c.readerState, 1)
	defer func() {
		atomic.StoreInt32(&c.readerState, 0)
	}()

	err := c.connection.SetReadDeadline(time.Now().Add(c.readTimeout))
	if err != nil {
		return err
	}

	c.connection.SetPongHandler(func(string) error {
		return c.connection.SetReadDeadline(time.Now().Add(c.readTimeout))
	})

	for {
		_, message, err := c.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				return err
			}

			return nil
		}

		c.responseMessages <- message
	}
}

func (c *WebsocketAPIClient) runWriter() error {
	atomic.StoreInt32(&c.writerState, 1)
	pingTicker := time.NewTicker((c.readTimeout * 9) / 10)
	defer func() {
		pingTicker.Stop()
		atomic.StoreInt32(&c.writerState, 0)
	}()

	for {
		select {
		case message, ok := <-c.requestMessages:
			if !ok {
				return nil
			}

			err := c.connection.SetWriteDeadline(time.Now().Add(c.writeTimeout))
			if err != nil {
				return err
			}

			err = c.connection.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return err
			}
		case <-pingTicker.C:
			err := c.connection.SetWriteDeadline(time.Now().Add(c.writeTimeout))
			if err != nil {
				return err
			}

			err = c.connection.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				return err
			}
		case <-c.quit:
			return nil
		}
	}
}

func (c *WebsocketAPIClient) sendAndWait(requestMessage []byte) ([]byte, error) {
	err := c.send(requestMessage)
	if err != nil {
		return nil, err
	}

	return c.wait()
}

func (c *WebsocketAPIClient) send(requestMessage []byte) error {
	select {
	case c.requestMessages <- requestMessage:
		return nil
	case <-time.After(c.writeTimeout):
		return fmt.Errorf("request sending timeout: %s", c.writeTimeout)
	}
}

func (c *WebsocketAPIClient) wait() ([]byte, error) {
	for {
		select {
		case message := <-c.responseMessages:
			if bytes.Contains(message, []byte("ms.remote.touchDisable")) ||
				bytes.Contains(message, []byte("ms.remote.touchEnable")) ||
				bytes.Contains(message, []byte("ms.remote.imeEnd")) {
				continue
			}

			return message, nil
		case <-time.After(c.readTimeout):
			return nil, fmt.Errorf("response waiting timeout: %s", c.writeTimeout)
		}
	}
}

func (c *WebsocketAPIClient) generateURL(token string) string {
	schema := "ws"
	if c.isSecure {
		schema = "wss"
	}

	return fmt.Sprintf(
		"%s://%s:%s/api/v2/channels/samsung.remote.control?name=%s&token=%s",
		schema,
		c.host,
		c.port,
		base64.URLEncoding.EncodeToString([]byte(c.clientID)),
		token,
	)
}
