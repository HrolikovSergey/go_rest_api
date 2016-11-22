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

var (
    settings Settings = Settings{};
    //DBConnectionString string
)


const mode = "debug" //debug | release | test

func Init(){
    os.Setenv("GIN_MODE", mode)
    file, err := os.Open("settings/settings.json")
    if (err != nil) {
        panic(err)
    }
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&settings)
    if err != nil {
        panic(err)
    }
    //DBConnectionString =
}

func Get() Settings {
    if settings == (Settings{}) {
        Init()
    }
    return settings
}