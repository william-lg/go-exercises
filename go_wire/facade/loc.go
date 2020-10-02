package facade

import (
	"github.com/google/wire"
	"github.com/william-lg/go-exercises/go_wire/config"
	"github.com/william-lg/go-exercises/go_wire/manager"
	"github.com/william-lg/go-exercises/go_wire/repo/mysql"
	"github.com/william-lg/go-exercises/go_wire/repo/mysql_mock"
	"github.com/william-lg/go-exercises/go_wire/service"
	"github.com/william-lg/go-exercises/go_wire/third_party/live"
	"github.com/william-lg/go-exercises/go_wire/third_party/sls"
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
	sls.AllThirdPartySet,
	config.AllConfigSet,
	live.AllThirdPartySet,
)

var LocationFacadeMockSet = wire.NewSet(
	NewLocationFacade,
	service.AllServiceSet,
	manager.AllManagerSet,
	mysql_mock.AllMysqlRepoMockSet,
	sls.AllThirdPartySet,
	config.AllConfigSet)

var LocationFacadeMockThirdParty = wire.NewSet(
	NewLocationFacade,
	service.AllServiceSet,
	manager.AllManagerSet,
	mysql_mock.AllMysqlRepoMockSet,
	sls.AllThirdPartyMockSet,
	config.AllConfigSet)
