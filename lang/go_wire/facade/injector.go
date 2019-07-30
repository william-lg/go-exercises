//+build wireinject

package facade

import "github.com/google/wire"

func InitLocationFacade() *LocationFacade {
	panic(wire.Build(LocationFacadeSet))

	return &LocationFacade{}
}

func InitWhsFacade() *WhsFacade {
	panic(wire.Build(WhsFacadeSet))

	return &WhsFacade{}
}
