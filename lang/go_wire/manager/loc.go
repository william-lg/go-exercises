package manager

import (
	"github.com/william-lg/go-exercises/lang/go_wire/repo"
	"github.com/william-lg/go-exercises/lang/go_wire/third_party"
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
	slsService third_party.SlsService,
	wmsService third_party.WmsService) *LocationManager {
	return &LocationManager{
		whsRepo:    whsRepo,
		locRepo:    locRepo,
		slsService: slsService,
		wmsService: wmsService,
	}
}
