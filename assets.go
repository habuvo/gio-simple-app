package main

import (
	"image"
	"image/color"
	"slices"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"

	"gioui.org/x/component"
)

var (
	colMain = color.NRGBA{A: 0xff, R: 0, G: 0, B: 0}
	colBack = color.NRGBA{A: 0xff, R: 111, G: 222, B: 222}
)

func makeTable(gtx layout.Context, th *material.Theme, headers []string, cells [][]string) layout.Dimensions {
	inset := layout.UniformInset(unit.Dp(2))

	// Configure a label styled to be a heading
	colHead := material.Body1(th, "")
	colHead.Alignment = text.Middle
	colHead.MaxLines = 1

	// Configure a label styled to be a cell
	cell := material.Body1(th, "")
	cell.Alignment = text.Middle
	cell.MaxLines = 1

	// Measure the height of a heading row.
	orig := gtx.Constraints
	gtx.Constraints.Min = image.Point{}
	macro := op.Record(gtx.Ops)
	dims := inset.Layout(gtx, colHead.Layout)
	_ = macro.Stop()
	gtx.Constraints = orig

	numCols := len(headers)

	return component.Table(th, &assetsGrid).Layout(gtx, len(cells), numCols,
		// Dimensioner func
		func(axis layout.Axis, index, constraint int) int {
			switch axis {
			case layout.Horizontal:
				minWidth := gtx.Dp(unit.Dp(50))
				return max(int(float32(constraint)/float32(numCols)), minWidth)
			default:
				return dims.Size.Y
			}
		},
		// Heading func
		func(gtx layout.Context, col int) layout.Dimensions {
			return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				colHead.Text = ""
				colName := headers[col]
				colHead.Text = colName
				colHead.Font.Weight = font.Bold
				return colHead.Layout(gtx)
			})
		},
		// Cell func
		func(gtx layout.Context, row, col int) layout.Dimensions {
			return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				cell.Text = ""
				cell.Color = colMain
				cell.Text = cells[row][col]
				cell.Alignment = text.Middle
				cell.Color = colMain

				paint.FillShape(gtx.Ops, colBack, clip.Rect{Max: gtx.Constraints.Max}.Op())

				return cell.Layout(gtx)
			})
		},
	)
}

func processOrder(name string, price, quantity int) {
	idx := slices.IndexFunc(myAssets, func(a asset) bool { return a.name == name })
	if idx == -1 {
		myAssets = append(myAssets, asset{name, quantity, price})

		return
	}

	myAssets[idx].quantity += quantity
}
