package view

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type LabelLayout struct {
	component.Tooltip
	material.LabelStyle
	State *component.TipArea
}

func (l LabelLayout) Layout(ctx layout.Context) layout.Dimensions {
	return l.State.Layout(ctx, l.Tooltip, l.LabelStyle.Layout)
}

type ImageLayout struct {
	image   image.Image
	imageOp paint.ImageOp
}

func (l ImageLayout) Layout(ctx layout.Context, width, height int) layout.Dimensions {
	if l.image != nil {
		l.imageOp.Add(ctx.Ops)
		paint.PaintOp{}.Add(ctx.Ops)
	} else {
		rect := clip.Rect{Max: image.Point{X: width, Y: height}}.Op()
		paint.FillShape(ctx.Ops, color.NRGBA{R: 200, G: 200, B: 200, A: 255}, rect)
	}
	return layout.Dimensions{Size: image.Point{X: width, Y: height}}
}
