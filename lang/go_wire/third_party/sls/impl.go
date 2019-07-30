package sls

import (
	"fmt"
	"github.com/william-lg/go-exercises/lang/go_wire/config"
)

type ServiceImpl struct {
	cfg *config.SlsConfig
}

func NewServiceImpl(cfg *config.SlsConfig) *ServiceImpl {
	return &ServiceImpl{
		cfg: cfg,
	}
}

//var ServiceImplSet = wire.NewSet(
//	NewServiceImpl,
//	wire.Bind(new(third_party.SlsService), new(ServiceImpl)))

func (ServiceImpl) CreateLogistic() {
	fmt.Println("create logistic")
}

func (ServiceImpl) CancelLogistic() {
	fmt.Println("Cancel logistic")
}
