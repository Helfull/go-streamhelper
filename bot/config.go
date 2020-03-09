package bot

import (
	"encoding/json"
	"io/ioutil"

	"github.com/helfull/go-streamhelper/util"
)

type Config struct {
	Server string `json:"server"`

	Debug bool `json:"debug"`

	Nickname string `json:"nickname"`
	Oauth    string `json:"Oauth"`
	Channel  string `json:"Channel"`
}

func GetConfig(configFile string) Config {

	raw, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic("File does not exists")
	}

	var config Config
	json.Unmarshal(raw, &config)

	config.loadEnvVars()
	return config
}

func (config *Config) loadEnvVars() (*Config, error) {
	util.LoadDotEnv()
	var err error
	config.Server, err = util.GetEnvStr("SH_SERVER", config.Server)
	if err != nil {
		return nil, err
	}
	config.Debug, err = util.GetEnvBool("SH_DEBUG", config.Debug)
	if err != nil {
		return nil, err
	}
	config.Channel, err = util.GetEnvStr("SH_CHANNEL", config.Channel)
	if err != nil {
		return nil, err
	}
	config.Oauth, err = util.GetEnvStr("SH_OAUTH", config.Oauth)
	if err != nil {
		return nil, err
	}
	config.Nickname, err = util.GetEnvStr("SH_NICKNAME", config.Nickname)
	if err != nil {
		return nil, err
	}

	return config, nil
}
