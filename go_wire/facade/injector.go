//+build wireinject

package facade

import "github.com/google/wire"

func InitLocationFacade() *LocationFacade {
	panic(wire.Build(LocationFacadeSet))

	return &LocationFacade{}
}

//func InitLocationFacadeMock() *LocationFacade {
//	panic(wire.Build(LocationFacadeMockSet))
//
//	return &LocationFacade{}
//}
//
//func InitLocationFacadeMockThirdParty() *LocationFacade {
//	panic(wire.Build(LocationFacadeMockThirdParty))
//
//	return nil
//}
//
//func InitWhsFacade() *WhsFacade {
//	panic(wire.Build(WhsFacadeSet))
//
//	return &WhsFacade{}
//}
