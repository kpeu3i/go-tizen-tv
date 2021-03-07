package samsung

import (
	"strconv"
)

type TVDevice map[string]string

func (d TVDevice) ID() string {
	return d["id"]
}

func (d TVDevice) IP() string {
	return d["ip"]
}

func (d TVDevice) Name() string {
	return d["name"]
}

func (d TVDevice) MAC() string {
	return d["wifiMac"]
}

func (d TVDevice) TokenAuthSupport() bool {
	if d["TokenAuthSupport"] == "" {
		return false
	}

	b, err := strconv.ParseBool(d["TokenAuthSupport"])
	if err != nil {
		return false
	}

	return b
}
