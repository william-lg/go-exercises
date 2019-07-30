package live

import (
	"fmt"
	"github.com/william-lg/go-exercises/lang/go_wire/config"
)

type WmsServiceImpl struct {
	slsCfg    *config.SlsConfig
	webapiCfg *config.WebapiConfig
}

func NewWmsServiceImpl(slsCg *config.SlsConfig, webapiCfg *config.WebapiConfig) *WmsServiceImpl {
	return &WmsServiceImpl{
		slsCfg:    slsCg,
		webapiCfg: webapiCfg,
	}
}

//var WmsServiceImplSet = wire.NewSet(
//	NewWmsServiceImpl,
//	wire.Bind(new(third_party.WmsService), new(WmsServiceImpl)))

func (WmsServiceImpl) CreateWmsOrder() {
	fmt.Println("create wms order")
}

func (WmsServiceImpl) CancelWmsOrder() {
	fmt.Println("cancel wms order")
}
