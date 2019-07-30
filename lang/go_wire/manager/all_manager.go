package manager

import "github.com/google/wire"

var AllManagerSet = wire.NewSet(
	NewWhsManager,
	NewLocationManager)
