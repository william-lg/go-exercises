package facade

import (
	"github.com/google/wire"
	"github.com/william-lg/go-exercises/go_wire/config"
	"github.com/william-lg/go-exercises/go_wire/manager"
	"github.com/william-lg/go-exercises/go_wire/repo/mysql"
	"github.com/william-lg/go-exercises/go_wire/service"
	"github.com/william-lg/go-exercises/go_wire/third_party/sls"
)

type WhsFacade struct {
	whsService      *service.WhsService
	locationService *service.LocationService
}

func NewWhsFacade(
	whsService *service.WhsService,
	locationService *service.LocationService) *WhsFacade {
	return &WhsFacade{
		whsService:      whsService,
		locationService: locationService,
	}
}

var WhsFacadeSet = wire.NewSet(
	NewWhsFacade,
	service.AllServiceSet,
	manager.AllManagerSet,
	mysql.AllMysqlRepoSet,
	sls.AllThirdPartySet,
	config.AllConfigSet)
