package mysql

import "github.com/google/wire"

var AllMysqlRepoSet = wire.NewSet(
	LocationRepoSet,
	WhsRepoSet)
