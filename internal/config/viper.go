package config

import (
	"flag"
	"log"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/spf13/viper"
)

func LoadConfig() error {
	var env string
	var confDir string
	flag.StringVar(&confDir, "confdir", "", "Specify local configuration directory") //configs
	flag.StringVar(&env, "env", "local", "Specify a env directory than default.")    //local or docker
	flag.Parse()

	path := determinePath(confDir, env)
	return viperSetting(path)
}

func viperSetting(configPath string) error {
	viper.SetConfigType(defaultConfigType)
	viper.AddConfigPath(configPath)
	viper.SetConfigName(serviceFileName)

	if err := viper.ReadInConfig(); err != nil {
		return errors.Errorf("viper.ReadInConfig error(%v)", err)
	}
	log.Println("viper.ConfigFileUsed OK:", viper.ConfigFileUsed())

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // export  SERVICE_NAME=test //Name Rule: Upper with '_'

	err := viper.Unmarshal(&configData) //bind struct
	if err != nil {
		return errors.Errorf("viper.Unmarshal:(%v)", err)
	}

	log.Println("data:", viper.Get("service.name"))        //TODO test case
	log.Println("Mongo:", viper.Get("Clients.Mongo.Host")) //map
	log.Println("Service.Name:", configData.Service.Name)
	return nil
}

const (
	defaultConfigDirectory = "./configs"
	defaultConfigType      = "toml"
	serviceFileName        = "service"
)

func determinePath(confDir, env string) string {
	path := confDir
	if len(path) == 0 {
		path = defaultConfigDirectory
	}

	path = filepath.Join(path, env)

	return path
}
