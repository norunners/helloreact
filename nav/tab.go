package nav

import (
	r "myitcv.io/react"
	"github.com/norunners/helloreact/app"
)

//go:generate reactGen

type TabDef struct {
	r.ComponentDef
}

type TabProps struct {
	Name string

	SetTitle app.SetTitle
}

func Tab(props TabProps) *TabDef {
	def := &TabDef{}
	r.BlessElement(def, props)
	return def
}

func (def *TabDef) Render() r.Element {
	props := def.Props()
	return r.Button(&r.ButtonProps{OnClick: def},
		r.S(props.Name),
	)
}

func (def *TabDef) OnClick(event *r.SyntheticMouseEvent) {
	props := def.Props()
	props.SetTitle.SetTitle(props.Name)
	event.PreventDefault()
}
