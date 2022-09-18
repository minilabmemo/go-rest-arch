package config

var configData *CofigDefinition

type CofigDefinition struct {
	Service ServiceInfo
	Log     LogInfo
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
