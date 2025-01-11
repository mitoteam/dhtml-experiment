package app

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/dhtmlbs"
	"github.com/mitoteam/dhtmlform"
)

func BuildDhtmlExperiment1() string {
	document := dhtml.NewHtmlDocument()

	document.
		Title("The Experiment!").
		Stylesheet("https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css")

	div := dhtml.Div().
		Id("test").
		Attribute("data-mt-test", "some attribute").
		//classes deduplication
		Class("border").Class("m-3").Class("p-3").Class("border").
		Text("some <escaped> text")

	document.Body().Append(div)

	document.Body().
		Append(
			dhtml.Div().Class("border").Class("t-3").Class("p-3").
				Text("multi").
				Append(
					dhtml.Span().Text("red").Class([]string{"border", "border-danger", "border-5"}),
				).
				Text("content"),
		).
		Append(
			dhtml.Div().Class("border p-3   m-3").
				Text("content").
				Text("only"),
		).
		Append(
			dhtml.Div().Class([]string{"border", "p-3", "m-3"}).
				Text("Icon test: "),
		)

	document.Body().Append(
		dhtmlbs.NewCard().
			Header(
				dhtmlbs.NewJustifiedLR().L("Card title text").R("Something right"),
			).
			Body("card body"),
	)

	return document.String()
}

var ExperimentFormHandler = dhtmlform.FormHandler{
	RenderF: func(formBody *dhtml.HtmlPiece, fd *dhtmlform.FormData) {
		formBody.Append("ExperimentFormHandler")
		formBody.Append(dhtmlform.NewTextarea("area").Default("def value\nmulti").Label("Label").Note("notes for <textarea>"))
		formBody.Append(dhtmlform.NewTextInput("txt").Default("def").Label("Label text").Note("notes for <input>"))

		formBody.Append(
			dhtml.Div().Append("Deeper").Class("mt-3 p-3 border").Append(
				dhtmlform.NewTextarea("area2").Default("def2").Label("Label2").Note("note2"),
			).Append(
				dhtml.Div().Append("And Deeper").Class("mt-3 p-3 border").Append(
					dhtmlform.NewTextarea("area3").Require().Label("Label3").Note("note3"),
				),
			),
		)

		formBody.Append(
			dhtmlform.NewCheckbox("cb1").Default(true).Label("Checkbox 1").Note("notes for checkbox"),
			dhtmlform.NewCheckbox("cb2").Label("Checkbox 2"),
			dhtmlform.NewPasswordInput("pwd").Label("Password").Default("ha"),
		)

		formBody.Append(dhtmlform.NewSubmitBtn())
	},

	ValidateF: func(fd *dhtmlform.FormData) {
		if v, ok := fd.GetValue("area2").(string); ok {
			if len(v) < 3 {
				fd.SetError("area2", "At least three characters expected")
			}
		}
	},

	SubmitF: func(fd *dhtmlform.FormData) {
		fd.SetRedirect("/")
	},
}
