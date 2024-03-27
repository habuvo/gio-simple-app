package main

import (
	"slices"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/event"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func handle(w *app.Window) error {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

	events := make(chan event.Event)
	acks := make(chan struct{})

	go func() {
		for {
			ev := w.NextEvent()
			events <- ev
			<-acks
			if _, ok := ev.(app.DestroyEvent); ok {
				return
			}
		}
	}()

	var ops op.Ops
	for {
		select {
		case e := <-events:
			switch e := e.(type) {
			case app.DestroyEvent:
				acks <- struct{}{}
				return e.Err
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				mainWidget(gtx, th)
				e.Frame(gtx.Ops)
			}
			acks <- struct{}{}
		case ut := <-updatesTicker:
			idx := slices.IndexFunc(myTickers, func(t ticker) bool { return ut.name == t.name })
			myTickers[idx].value = ut.value
			myTickers[idx].diff = ut.diff
			w.Invalidate()
		}
	}
}
