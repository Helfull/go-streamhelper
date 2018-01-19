package bot

import (
    "encoding/json"
    "io/ioutil"
)

type Config struct {
    Server string `json:"server"`

    Debug bool `json:"debug"`

    Nickname string `json:"nickname"`
    Oauth string `json:"Oauth"`
    Channel string `json:"Channel"`
}

func GetConfig(configFile string) Config {
    raw, err := ioutil.ReadFile(configFile)
    if err != nil {
        panic("File does not exists")
    }

    var config Config
    json.Unmarshal(raw, &config)

    return config
}