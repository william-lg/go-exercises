package sls

import (
	"github.com/google/wire"
)

var AllThirdPartySet = wire.NewSet(
	ServiceImplSet)

var AllThirdPartyMockSet = wire.NewSet(
	NewServiceImpl)
