package third_party

import (
	"github.com/google/wire"
	"github.com/william-lg/go-exercises/lang/go_wire/third_party/live"
	"github.com/william-lg/go-exercises/lang/go_wire/third_party/sls"
)

var AllThirdPartySet = wire.NewSet(
	sls.NewServiceImpl,
	live.NewWebapiServiceImpl,
	live.NewWmsServiceImpl)
