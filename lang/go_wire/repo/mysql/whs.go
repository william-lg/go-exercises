package mysql

import (
	"fmt"
	"github.com/google/wire"
	"github.com/william-lg/go-exercises/lang/go_wire/repo"
)

type Whs struct {
}

func NewWhsRepo() *Whs {
	return &Whs{}
}

var WhsRepoSet = wire.NewSet(
	NewWhsRepo,
	wire.Bind(new(repo.WhsIfc), new(*Whs)))

func (Whs) JustPrintWhs() {
	fmt.Println("Whs Repo Print")
}
