package center

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	Addr string `json:"addr"`
}

var c config

func load() error {
	buf, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, &c)
}
