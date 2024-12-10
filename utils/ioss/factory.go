package oss

import "crmeb_go/config"

type Factory interface {
	Create(c config.Conf) (Client, error)
}

var factoryMap = map[string]Factory{}

func registerFactory(provider string, factory Factory) {
	factoryMap[provider] = factory
}

func getFactory(provider string) (Factory, bool) {
	factory, ok := factoryMap[provider]
	return factory, ok
}
