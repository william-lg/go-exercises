//+build wireinject

package service

import "github.com/google/wire"

func InitLocationService() *LocationService {
	panic(wire.Build(locationServiceSet))

	return &LocationService{}
}

func InitWhsService() *WhsService {
	panic(wire.Build(whsServiceSet))

	return &WhsService{}
}
