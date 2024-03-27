package main

import (
	"fmt"
	"slices"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type asset struct {
	name      string
	quantity  int
	basePrice int
}

type ticker struct {
	button widget.Clickable
	name   string
	value  int
	diff   int
}

func makeCell(th *material.Theme, gtx layout.Context, name string, value, diff int) layout.Dimensions {
	clr := greenColor
	switch {
	case diff < 0:
		clr = redColor
	case len(name) == 0:
		clr = grayColor
	}

	var cell layout.Widget
	if len(name) == 0 {
		cell = material.Body1(th, ".").Layout
		//paint.FillShape(gtx.Ops, clr, clip.Rect{Max: gtx.Constraints.Max}.Op())
	} else {
		var cellButton *widget.Clickable
		idx := slices.IndexFunc(myTickers, func(t ticker) bool { return t.name == name })
		if idx >= 0 {
			cellButton = &myTickers[idx].button
		} else {
			cellButton = &widget.Clickable{}
		}

		b := material.Button(th, cellButton, fmt.Sprintf("%s\n%d (%d)", name, value, diff))
		b.Color = clr
		//b.Background = blackColor
		b.TextSize = 12
		// b.Inset = layout.Inset{
		// 	Top: 4, Bottom: 4,
		// 	Left: 4, Right: 4,
		// }
		cell = b.Layout
	}

	return layout.Center.Layout(gtx, cell)
}

var myTickers = []ticker{
	{widget.Clickable{}, "BTC", 1000, 12},
	{widget.Clickable{}, "DODGE", 2435, 211},
	{widget.Clickable{}, "USDT", 3333, 323},
}

func getValuesByIndex(idx int) (name string, value, diff int) {
	if idx < 0 || idx >= len(myTickers) {
		return
	}

	name = myTickers[idx].name
	value = myTickers[idx].value
	diff = myTickers[idx].diff

	return
}

var myAssets = []asset{
	{"BTC", 1000, 100},
	{"DODGE", 2000, 200},
	{"USDT", 3000, 300},
}
