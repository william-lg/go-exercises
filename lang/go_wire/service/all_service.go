package service

import "github.com/google/wire"

var AllServiceSet = wire.NewSet(
	NewLocationService,
	NewWhsService)
