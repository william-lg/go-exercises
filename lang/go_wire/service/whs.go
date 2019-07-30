package service

import (
	"github.com/google/wire"
	"github.com/william-lg/go-exercises/lang/go_wire/config"
	"github.com/william-lg/go-exercises/lang/go_wire/manager"
	"github.com/william-lg/go-exercises/lang/go_wire/repo/mysql"
	"github.com/william-lg/go-exercises/lang/go_wire/third_party"
)

type WhsService struct {
	whsManager      *manager.WhsManager
	locationService *LocationService
}

func NewWhsService(
	whsManager *manager.WhsManager,
	locationService *LocationService) *WhsService {
	return &WhsService{
		whsManager:      whsManager,
		locationService: locationService,
	}
}

var whsServiceSet = wire.NewSet(
	NewWhsService,
	NewLocationService,
	manager.AllManagerSet,
	mysql.AllMysqlRepoSet,
	third_party.AllThirdPartySet,
	config.AllConfigSet)
