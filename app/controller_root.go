package app

import (
	"github.com/mitoteam/mbr"
)

type RootController struct {
	mbr.ControllerBase
}

var RootCtl *RootController

func init() {
	RootCtl = &RootController{}

	//using chi middlewares
	//RootCtl.With(middleware.Recoverer)
}

func (c *RootController) Home() mbr.Route {
	return mbr.Route{
		PathPattern: "/",
		HandleF: func(ctx *mbr.MbrContext) any {
			return "Hello"
		},
	}
}
