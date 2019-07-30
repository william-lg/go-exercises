package facade

import (
	"github.com/google/wire"
	"github.com/william-lg/go-exercises/lang/go_wire/config"
	"github.com/william-lg/go-exercises/lang/go_wire/manager"
	"github.com/william-lg/go-exercises/lang/go_wire/repo/mysql"
	"github.com/william-lg/go-exercises/lang/go_wire/service"
	"github.com/william-lg/go-exercises/lang/go_wire/third_party"
)

type LocationFacade struct {
	locationService *service.LocationService
}

func NewLocationFacade(locationService *service.LocationService) *LocationFacade {
	return &LocationFacade{
		locationService: locationService,
	}
}

var LocationFacadeSet = wire.NewSet(
	NewLocationFacade,
	service.AllServiceSet,
	manager.AllManagerSet,
	mysql.AllMysqlRepoSet,
	third_party.AllThirdPartySet,
	config.AllConfigSet)
