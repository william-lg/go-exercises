// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package facade

import (
	"github.com/william-lg/go-exercises/lang/go_wire/config"
	"github.com/william-lg/go-exercises/lang/go_wire/manager"
	"github.com/william-lg/go-exercises/lang/go_wire/repo/mysql"
	"github.com/william-lg/go-exercises/lang/go_wire/service"
	"github.com/william-lg/go-exercises/lang/go_wire/third_party/live"
	"github.com/william-lg/go-exercises/lang/go_wire/third_party/sls"
)

// Injectors from injector.go:

func InitLocationFacade() *LocationFacade {
	whs := mysql.NewWhsRepo()
	location := mysql.NewLocationRepo()
	slsConfig := config.NewSlsConfig()
	serviceImpl := sls.NewServiceImpl(slsConfig)
	webapiConfig := config.NewWebapiConfig()
	wmsServiceImpl := live.NewWmsServiceImpl(slsConfig, webapiConfig)
	locationManager := manager.NewLocationManager(whs, location, serviceImpl, wmsServiceImpl)
	orderManager := manager.NewOrderManager()
	locationService := service.NewLocationService(locationManager, orderManager)
	locationFacade := NewLocationFacade(locationService)
	return locationFacade
}
