package app

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/mbr"
)

type RootController struct {
	mbr.ControllerBase
}

var RootCtl *RootController

func init() {
	RootCtl = &RootController{}
	RootCtl.With(RootMiddleware)

	//using chi middlewares
	//RootCtl.With(middleware.Recoverer)
}

func (c *RootController) Home() mbr.Route {
	return mbr.Route{
		PathPattern: "/",
		HandleF: func(ctx *mbr.MbrContext) any {
			out := dhtml.Piece("Hello")

			out.Append(
				dhtml.Div().Append(
					dhtml.NewLink(mbr.Url(c.Experiment)).Label("dhtml experiment1"),
				),
			)

			return out.String()
		},
	}
}

func (c *RootController) Experiment() mbr.Route {
	return mbr.Route{
		PathPattern: "/dhtml/experiment1",
		HandleF:     func(ctx *mbr.MbrContext) any { return BuildDhtmlExperiment1() },
	}
}
