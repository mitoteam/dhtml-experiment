package app

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/dhtmlbs"
	"github.com/mitoteam/dhtmlform"
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
			document := dhtml.NewHtmlDocument()

			document.
				Title("Dhtml Experiment").
				Stylesheet("https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css")

			mainOut := dhtml.Div().Id("main").Class("m-5")

			mainOut.Append(
				dhtml.RenderValue(
					"Source code",
					dhtml.NewLink("https://github.com/mitoteam/dhtml-experiment").Label("Github").Target("_blank"),
				).Class("p-3 border"),
			)

			btnPanel := dhtml.Div().Class("mt-3", "border", "p-3").Append(
				dhtmlbs.NewBtn().Href(mbr.Url(c.Experiment1)).Label("dhtml experiment1").Class("btn-success"),
				dhtmlbs.NewBtn().Href(mbr.Url(c.FormExperiment1)).Label("form experiment1"),
			)

			mainOut.Append(btnPanel)

			document.Body().Append(mainOut)

			return document.String()
		},
	}
}

func (c *RootController) Experiment1() mbr.Route {
	return mbr.Route{
		PathPattern: "/dhtml/experiment1",
		HandleF:     func(ctx *mbr.MbrContext) any { return BuildDhtmlExperiment1() },
	}
}

func (c *RootController) FormExperiment1() mbr.Route {
	return mbr.Route{
		PathPattern: "/dhtmlform/experiment1",
		HandleF: func(ctx *mbr.MbrContext) any {
			fc := dhtmlform.NewFormContext(ctx.Writer(), ctx.Request())

			// some parameters
			fc.SetParam("MbrContext", ctx)
			fc.SetParam("MyParameter", "we-ha!")

			return formHandlerExperiment1.Render(fc)
		},
	}
}
