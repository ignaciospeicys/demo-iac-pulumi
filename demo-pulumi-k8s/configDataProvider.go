package main

type ConfigDataProvider interface {
	GetConfigData() (map[string]string, error)
}
