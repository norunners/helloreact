package app

import (
	"github.com/norunners/helloreact/nav"
	r "myitcv.io/react"
)

//go:generate reactGen

type AppDef struct {
	r.ComponentDef
}

type AppState struct {
	title string
}

func App() *AppDef {
	def := &AppDef{}
	r.BlessElement(def, nil)
	return def
}

func (def *AppDef) ComponentWillMount() {
	def.SetTitle("Hello World!")
}

func (def *AppDef) Render() r.Element {
	state := def.State()
	return r.Div(nil,
		r.H1(nil,
			r.S(state.title),
		),
		r.P(nil,
			r.S("This is my first GopherJS React App."),
		),
		nav.Tab(nav.TabProps{Name: "Create", SetTitle: def}),
		nav.Tab(nav.TabProps{Name: "Read", SetTitle: def}),
		nav.Tab(nav.TabProps{Name: "Update", SetTitle: def}),
		nav.Tab(nav.TabProps{Name: "Delete", SetTitle: def}),
	)
}

type SetTitle interface {
	SetTitle(title string)
}

func (def *AppDef) SetTitle(title string) {
	state := def.State()
	state.title = title
	def.SetState(state)
}
