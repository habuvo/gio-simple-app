package main

import (
	"strconv"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

var (
	makeGrid = func(th *material.Theme) layout.Widget {
		return func(gtx layout.Context) layout.Dimensions {
			return component.Grid(th, &assetsGrid).Layout(gtx, gridX, gridY,
				func(axis layout.Axis, index, constraint int) int {
					return gtx.Dp(cellSize)
				},
				func(gtx layout.Context, row, col int) layout.Dimensions {
					name, value, diff := getValuesByIndex((row * gridY) + col)
					return makeCell(th, gtx, name, value, diff)
				})
		}
	}

	assetsTable = func(th *material.Theme) layout.Widget {
		return func(gtx layout.Context) layout.Dimensions {
			cells := make([][]string, len(myAssets))
			for i, t := range myAssets {
				cells[i] = append(cells[i],
					t.name, strconv.Itoa(t.basePrice),
					strconv.Itoa(t.quantity),
					strconv.Itoa(t.quantity*t.basePrice))
			}

			return makeTable(gtx, th, []string{"name", "base price", "quantity", "total"}, cells)
		}
	}

	adviser = func(th *material.Theme) layout.Widget {
		return func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Min.Y = gtx.Dp(unit.Dp(50))
				gtx.Constraints.Min.X = gtx.Dp(unit.Dp(400))
				adviseEditor.SetText("here will be a wise advice")
				adviseEditor.Alignment = text.Middle
				return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return material.Editor(th, &adviseEditor, "advise").Layout(gtx)
				})
			})
		}
	}
)
