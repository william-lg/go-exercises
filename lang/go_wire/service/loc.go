package service

import (
	"github.com/google/wire"
	"github.com/william-lg/go-exercises/lang/go_wire/config"
	"github.com/william-lg/go-exercises/lang/go_wire/manager"
	"github.com/william-lg/go-exercises/lang/go_wire/repo/mysql"
	"github.com/william-lg/go-exercises/lang/go_wire/third_party"
)

type LocationService struct {
	locationManager *manager.LocationManager
}

func NewLocationService(
	locationManager *manager.LocationManager) *LocationService {
	return &LocationService{
		locationManager: locationManager,
	}
}

var locationServiceSet = wire.NewSet(
	NewLocationService,
	manager.AllManagerSet,
	mysql.AllMysqlRepoSet,
	third_party.AllThirdPartySet,
	config.AllConfigSet)
