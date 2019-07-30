package manager

import (
	"github.com/william-lg/go-exercises/lang/go_wire/repo"
	"github.com/william-lg/go-exercises/lang/go_wire/third_party"
	"github.com/william-lg/go-exercises/lang/go_wire/third_party/live"
	"github.com/william-lg/go-exercises/lang/go_wire/third_party/sls"
)

type LocationManager struct {
	whsRepo    repo.WhsIfc
	locRepo    repo.LocationIfc
	slsService third_party.SlsService
	wmsService third_party.WmsService
}

func NewLocationManager(
	whsRepo repo.WhsIfc,
	locRepo repo.LocationIfc,
	slsService *sls.ServiceImpl,
	wmsService *live.WmsServiceImpl) *LocationManager {
	return &LocationManager{
		whsRepo:    whsRepo,
		locRepo:    locRepo,
		slsService: slsService,
		wmsService: wmsService,
	}
}
