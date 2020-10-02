package mysql

import (
	"fmt"
	"github.com/google/wire"
	"github.com/william-lg/go-exercises/go_wire/repo"
)

type Location struct {
}

func NewLocationRepo() *Location {
	return &Location{}
}

var LocationRepoSet = wire.NewSet(
	NewLocationRepo,
	wire.Bind(new(repo.LocationIfc), new(*Location)))

func (loc Location) JustPrintLocation() {
	fmt.Println("Location Repo Print")
}
