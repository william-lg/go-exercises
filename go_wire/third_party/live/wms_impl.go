package live

import (
	"fmt"
	"github.com/google/wire"
	"github.com/william-lg/go-exercises/go_wire/config"
	"github.com/william-lg/go-exercises/go_wire/third_party"
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

var WmsServiceImplSet = wire.NewSet(
	NewWmsServiceImpl,
	wire.Bind(new(third_party.WmsService), new(*WmsServiceImpl)))

func (WmsServiceImpl) CreateWmsOrder() {
	fmt.Println("create wms order")
}

func (WmsServiceImpl) CancelWmsOrder() {
	fmt.Println("cancel wms order")
}
