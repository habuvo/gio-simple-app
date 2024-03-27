package main

import (
	"image/color"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

var (
	amountInput  = widget.Editor{SingleLine: true, Submit: true}
	assetInput   = widget.Editor{SingleLine: true, Submit: true}
	tickerInput  = widget.Editor{SingleLine: true, Submit: true}
	orderButton  = new(widget.Clickable)
	deleteButton = new(widget.Clickable)
	cancelButton = new(widget.Clickable)

	tickerButton = new(widget.Clickable)
	geminiButton = new(widget.Clickable)
	bidaskGroup  = new(widget.Enum)

	orderAllowed = false
	adviseEditor widget.Editor

	list = &widget.List{
		List: layout.List{
			Axis: layout.Vertical,
		},
	}

	greenColor = color.NRGBA{A: 0xff, R: 0x00, G: 0xff, B: 0x00}
	redColor   = color.NRGBA{A: 0xff, R: 0xff, G: 0x00, B: 0x00}
	grayColor  = color.NRGBA{A: 0xff, R: 203, G: 205, B: 207}
	blackColor = color.NRGBA{A: 0xff, R: 0x00, G: 0x00, B: 0x00}

	in     = layout.UniformInset(unit.Dp(8))
	border = widget.Border{Color: color.NRGBA{A: 0xff}, CornerRadius: unit.Dp(8), Width: unit.Dp(2)}

	orderAsset string
	orderPrice int

	assetsGrid  component.GridState
	tickersGrid component.GridState

	cellSize = unit.Dp(70)
)

const (
	gridX = 4
	gridY = 6
)

func mainWidget(gtx layout.Context, th *material.Theme) layout.Dimensions {
	if orderButton.Clicked(gtx) {
		orderClicked()
	}

	if tickerButton.Clicked(gtx) {
		tickerClicked()
	}

	if deleteButton.Clicked(gtx) {
		deleteClicked()
	}

	if cancelButton.Clicked(gtx) {
		cancelClicked()
	}

	processMyTickersButtons(gtx)

	widgets := []layout.Widget{
		func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Axis:      layout.Vertical,
				Alignment: layout.Middle,
				Spacing:   layout.SpaceAround}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						if !orderAllowed {
							gtx = gtx.Disabled()
						}
						return layout.Flex{}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return material.Button(th, orderButton, "make order").Layout(gtx)
								})
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return material.Button(th, deleteButton, "delete ticker").Layout(gtx)
								})
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return material.Button(th, cancelButton, "cancel").Layout(gtx)
								})
							}),
							layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return layout.Flex{}.
										Layout(gtx,
											layout.Rigid(material.RadioButton(th, bidaskGroup, "bid", "bid").Layout),
											layout.Rigid(material.RadioButton(th, bidaskGroup, "ask", "ask").Layout),
											layout.Flexed(10, func(gtx layout.Context) layout.Dimensions {
												l := material.Label(th, 18, orderAsset)
												l.Alignment = text.Alignment(layout.Middle)
												return l.Layout(gtx)
											}),
											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												e := material.Editor(th, &amountInput, "amount")
												e.Font.Style = font.Italic
												border := widget.Border{Color: color.NRGBA{A: 0xff}, CornerRadius: unit.Dp(8), Width: unit.Dp(2)}
												return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
													return layout.UniformInset(unit.Dp(8)).Layout(gtx, e.Layout)
												})
											}),
										)
								})
							}),
						)
					})
				}),
				layout.Rigid(
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Alignment: layout.Baseline}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return material.Button(th, tickerButton, "add ticker").Layout(gtx)
								})
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return layout.Flex{Alignment: layout.Alignment(layout.SpaceEvenly)}.
										Layout(gtx,
											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												e := material.Editor(th, &tickerInput, "ticker")
												e.Font.Style = font.Italic
												border := widget.Border{Color: color.NRGBA{A: 0xff}, CornerRadius: unit.Dp(8), Width: unit.Dp(2)}
												return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
													return layout.UniformInset(unit.Dp(8)).Layout(gtx, e.Layout)
												})
											}),
										)
								})
							}),
						)
					})
				}),
			)
		},
		func(gtx layout.Context) layout.Dimensions {
			return material.Label(th, 24, "your watch list").Layout(gtx)
		},
		func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, makeGrid(th))
		},
		func(gtx layout.Context) layout.Dimensions {
			return material.Label(th, 24, "your assets").Layout(gtx)
		},
		assetsTable(th),
		func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Height: unit.Dp(10)}.Layout(gtx)
		},
		func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Button(th, geminiButton, "ask Gemini for advice").Layout(gtx)
			})
		},
		adviser(th),
	}

	return material.List(th, list).Layout(gtx, len(widgets),
		func(gtx layout.Context, i int) layout.Dimensions {
			return layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
		},
	)
}
