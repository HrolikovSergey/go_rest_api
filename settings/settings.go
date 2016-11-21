package settings

import (
    "os"
    "encoding/json"
)

type(
    Settings struct {
        DbHost     string
        DbName     string
        DbUser     string
        DbPassword string
        DbCharset  string

        ServerPort string

        Salt string
        AuthTokenExpire int64
        PrivateKeyPath string
        PublicKeyPath string
    }
)

var settings Settings = Settings{};


func Init(){
    file, err := os.Open("settings/settings.json")
    if (err != nil) {
        panic(err)
    }
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&settings)
    if err != nil {
        panic(err)
    }
}

func Get() Settings {
    if settings == (Settings{}) {
        Init()
    }
    return settings
}