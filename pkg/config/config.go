package config

import (
	"errors"
	"os"

        flag "github.com/spf13/pflag"
        "github.com/spf13/viper"

        "github.com/BurntSushi/toml"
)


func ParseFlags() {
        viper.AutomaticEnv()
        flag.String("DBAddr", "", "Backend Database host:port; overrides DBADDR")
        flag.String("DBUser", "", "Backend Database username; overrides DBUSER")
        flag.String("DBPass", "", "Backend Database password; overrides DBPASS")
        flag.String("serviceFile", "services.toml", "TOML file to read service definitions from; overrides SERVICEFILE")
        flag.Bool("help", false, "print this help and exit")
        flag.Parse()
        viper.BindPFlags(flag.CommandLine)
}


func LoadConfig(configFile string) (*Config, error) {
    if _, err := os.Stat(configFile); os.IsNotExist(err) {
        return nil, errors.New("Config file does not exist.")
    } else if err != nil {
        return nil, err
    }

    var conf Config
    if _, err := toml.DecodeFile(configFile, &conf); err != nil {
        return nil, err
    }

    return &conf, nil
}
