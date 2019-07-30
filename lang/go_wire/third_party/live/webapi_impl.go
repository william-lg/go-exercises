package live

import (
	"fmt"
	"github.com/william-lg/go-exercises/lang/go_wire/config"
)

type WebapiServiceImpl struct {
	slsCfg    *config.SlsConfig
	webapiCfg *config.WebapiConfig
}

func NewWebapiServiceImpl(slsCg *config.SlsConfig, webapiCfg *config.WebapiConfig) *WebapiServiceImpl {
	return &WebapiServiceImpl{
		slsCfg:    slsCg,
		webapiCfg: webapiCfg,
	}
}

//var WebapiServiceImplSet = wire.NewSet(
//	NewWebapiServiceImpl,
//	wire.Bind(new(third_party.WebapiService), new(WebapiServiceImpl)))

func (WebapiServiceImpl) CreateOrder() {
	fmt.Println("create order")
}

func (WebapiServiceImpl) CancelOrder() {
	fmt.Println("cancel order")
}
