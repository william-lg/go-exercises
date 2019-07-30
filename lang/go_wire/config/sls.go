package config

import "github.com/google/wire"

type SlsConfig struct {
	Host string
	Url  string
}

func NewSlsConfig() *SlsConfig {
	return &SlsConfig{
		Host: "localhost:8080",
		Url:  "/hello_world",
	}
}

type WebapiConfig struct {
	Host string
	Url  string
}

func NewWebapiConfig() *WebapiConfig {
	return &WebapiConfig{
		Host: "localhost:8080",
		Url:  "/hello_world",
	}
}

type WmsConfig struct {
	Host string
	Url  string
}

func NewWmsConfig() *WmsConfig {
	return &WmsConfig{
		Host: "localhost:8080",
		Url:  "/hello_world",
	}
}

var AllConfigSet = wire.NewSet(
	NewSlsConfig,
	NewWebapiConfig,
	NewWmsConfig)
