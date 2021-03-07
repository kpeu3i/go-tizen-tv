package tizenapi

type ConnectResponseMessage struct {
	Event string `json:"event"`
	Data  struct {
		ID      string `json:"id"`
		Token   string `json:"token"`
		Clients []struct {
			ID          string  `json:"id"`
			ConnectTime float64 `json:"connectTime"`
			DeviceName  string  `json:"deviceName"`
			IsHost      bool    `json:"isHost"`
			Attributes  struct {
				Name string `json:"name"`
			} `json:"attributes"`
		} `json:"clients"`
	} `json:"data"`
}

type GetAppsRequestMessage struct {
	Method string `json:"method"`
	Params struct {
		Event string `json:"event"`
		To    string `json:"to"`
	} `json:"params"`
}

type GetAppsResponseMessage struct {
	Event string `json:"event"`
	From  string `json:"from"`
	Data  struct {
		Data []struct {
			AppId   string `json:"appId"`
			AppType int    `json:"app_type"`
			Icon    string `json:"icon"`
			Name    string `json:"name"`
		} `json:"data"`
	} `json:"data"`
}

type OpenAppRequestMessage struct {
	Method string `json:"method"`
	Params struct {
		Event string `json:"event"`
		To    string `json:"to"`
		Data  struct {
			ActionType string `json:"action_type"`
			AppID      string `json:"appId"`
			MetaTag    string `json:"metaTag"`
		} `json:"data"`
	} `json:"params"`
}

type SendKeyRequestMessage struct {
	Method string `json:"method"`
	Params struct {
		Cmd          string `json:"Cmd"`
		DataOfCmd    string `json:"DataOfCmd"`
		Option       string `json:"Option"`
		TypeOfRemote string `json:"TypeOfRemote"`
	} `json:"params"`
}
