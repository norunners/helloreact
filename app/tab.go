package app

import (
	r "myitcv.io/react"
)

//go:generate reactGen

type TabDef struct {
	r.ComponentDef

	setTitle SetTitle
}

type TabProps struct {
	Name string
}

func Tab(props TabProps, setTitle SetTitle) *TabDef {
	def := &TabDef{setTitle: setTitle}
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
	def.setTitle.SetTitle(props.Name)
	event.PreventDefault()
}
