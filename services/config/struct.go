package config

import (
	"fmt"
	"strings"
)

var ConfigData *CofigDefinition

type CofigDefinition struct {
	Service ServiceInfo
	Log     LogInfo
	Clients map[string]ClientInfo
}

type ServiceInfo struct {
	Name       string
	Port       int
	StartupMsg string
}

type LogInfo struct {
	Level string
	File  string
}

type ClientInfo struct {
	Host     string
	Port     int
	Protocol string
	Username string
	Password string
}

func (c ClientInfo) Url(apiRoute string) string {

	if !strings.HasPrefix(apiRoute, "/") {
		apiRoute = fmt.Sprintf("/%s", apiRoute)
	}

	url := fmt.Sprintf("%s://%s:%d%s", c.Protocol, c.Host, c.Port, apiRoute)

	return url
}
