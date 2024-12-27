package view

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"image/color"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type UI struct {
	Theme *material.Theme
	grid  component.GridState
}

func NewUI() *UI {
	ui := new(UI)
	ui.Theme = material.NewTheme()
	ui.Theme.Face = "JetBrains Mono"
	return ui
}

func (ui *UI) Run(window *app.Window) error {
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			ui.Layout(gtx)

			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			return e.Err
		}
	}
}

func (ui *UI) Layout(gtx C) D {
	const sideLength int = 8
	const cellSize int = 90

	clickers := make([]widget.Clickable, 0)
	for i := 0; i < sideLength*sideLength; i++ {
		clickers = append(clickers, widget.Clickable{})
	}

	return component.Grid(ui.Theme, &ui.grid).Layout(gtx, sideLength, sideLength,
		func(axis layout.Axis, index, constraint int) int {
			return gtx.Dp(unit.Dp(cellSize))
		},
		func(gtx C, row, col int) D {
			clk := &clickers[row*sideLength+col]
			btn := material.Button(ui.Theme, clk, fmt.Sprintf("R%d C%d", row, col))
			color := color.NRGBA{
				R: uint8(255 / sideLength * row),
				G: uint8(255 / sideLength * col),
				B: uint8(255 * row * col / (sideLength * sideLength)),
				A: 255,
			}
			btn.Background = color
			btn.CornerRadius = 0
			return btn.Layout(gtx)
		})
}
