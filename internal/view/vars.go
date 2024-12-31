package view

import (
	"gioui.org/op/paint"
	"gioui.org/unit"
	"image"
)

var (
	dp1  = unit.Dp(1)
	dp8  = unit.Dp(8)
	dp16 = unit.Dp(16)
	dp24 = unit.Dp(24)
	dp40 = unit.Dp(40)
)

type Video struct {
	Id        string
	Name      string
	Tags      string
	Path      string
	Size      int64
	Cover     []byte
	Format    string
	Duration  int64
	CreatedAt int64
	UpdatedAt int64

	CoverImage   image.Image
	CoverImageOp paint.ImageOp
}
