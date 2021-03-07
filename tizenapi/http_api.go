package tizenapi

type GetAppResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Running bool   `json:"running"`
	Visible bool   `json:"visible"`
	Version string `json:"version"`
}

type GetInfoResponse struct {
	ID        string            `json:"id"`
	Type      string            `json:"type"`
	Name      string            `json:"name"`
	Version   string            `json:"version"`
	URI       string            `json:"uri"`
	Remote    string            `json:"remote"`
	Device    map[string]string `json:"device"`
	IsSupport string            `json:"isSupport"`
}
