package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type config struct {
	Site struct {
		Name    string `toml:"name"`
		BaseURL string `toml:"base_url"`
		Port    int    `toml:"port"`
	} `toml:"site"`

	MySQL struct {
		User     string `toml:"user"`
		Password string `toml:"password"`
		Addr     string `toml:"addr"`
		Name     string `toml:"name"`
	} `toml:"mysql"`

	TgBot struct {
		Token         string `toml:"token"`
		TargetChannel string `toml:"targetChannel"`
		//TargetUser    int64  `toml:"targetUser"`
	}
}

var conf *config

func Load() error {
	c := config{}

	_, err := toml.DecodeFile("./config/kyoto.toml", &c)
	if err != nil {
		return errors.Wrap(err, "decode config file")
	}

	conf = &c
	return nil
}

// Get returns the config struct.
func Get() *config {
	return conf
}
