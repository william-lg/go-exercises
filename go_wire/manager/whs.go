package manager

import (
	"github.com/william-lg/go-exercises/go_wire/repo"
	"github.com/william-lg/go-exercises/go_wire/third_party"
	"github.com/william-lg/go-exercises/go_wire/third_party/sls"
)

type WhsManager struct {
	whsRepo    repo.WhsIfc
	slsService third_party.SlsService
}

func NewWhsManager(
	whsRepo repo.WhsIfc,
	slsService *sls.ServiceImpl) *WhsManager {
	return &WhsManager{
		whsRepo:    whsRepo,
		slsService: slsService,
	}
}
