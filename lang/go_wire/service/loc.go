package service

import (
	"github.com/google/wire"
	"github.com/william-lg/go-exercises/lang/go_wire/config"
	"github.com/william-lg/go-exercises/lang/go_wire/manager"
	"github.com/william-lg/go-exercises/lang/go_wire/repo/mysql"
	"github.com/william-lg/go-exercises/lang/go_wire/third_party/sls"
)

type LocationService struct {
	locationManager *manager.LocationManager
	orderManager    *manager.OrderManager
}

func NewLocationService(
	locationManager *manager.LocationManager,
	orderManager *manager.OrderManager) *LocationService {
	return &LocationService{
		locationManager: locationManager,
		orderManager:    orderManager,
	}
}

var locationServiceSet = wire.NewSet(
	NewLocationService,
	manager.AllManagerSet,
	mysql.AllMysqlRepoSet,
	sls.AllThirdPartySet,
	config.AllConfigSet)
