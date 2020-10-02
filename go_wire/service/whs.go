package service

import (
	"github.com/google/wire"
	"github.com/william-lg/go-exercises/go_wire/config"
	"github.com/william-lg/go-exercises/go_wire/manager"
	"github.com/william-lg/go-exercises/go_wire/repo/mysql"
	"github.com/william-lg/go-exercises/go_wire/third_party/sls"
)

type WhsService struct {
	whsManager      *manager.WhsManager
	locationService *LocationService
	orderManager    *manager.OrderManager
}

func NewWhsService(
	whsManager *manager.WhsManager,
	locationService *LocationService,
	orderManager *manager.OrderManager) *WhsService {
	return &WhsService{
		whsManager:      whsManager,
		locationService: locationService,
		orderManager:    orderManager,
	}
}

var whsServiceSet = wire.NewSet(
	NewWhsService,
	NewLocationService,
	manager.AllManagerSet,
	mysql.AllMysqlRepoSet,
	sls.AllThirdPartySet,
	config.AllConfigSet)
