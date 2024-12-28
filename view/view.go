package view

import (
	"fmt"
	"image"
	"image/color"

	"gioui.org/op/clip"
	"gioui.org/op/paint"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type UI struct {
	Theme *material.Theme

	search       widget.Editor
	searchButton widget.Clickable
	grid         component.GridState
}

func NewUI() *UI {
	ui := new(UI)
	ui.Theme = material.NewTheme()
	ui.search = widget.Editor{
		SingleLine: true,
	}
	return ui
}

func (ui *UI) Run(window *app.Window) error {
	ops := new(op.Ops)
	for {
		switch e := window.Event().(type) {
		case app.FrameEvent:
			c := app.NewContext(ops, e)
			ui.Layout(c)
			e.Frame(c.Ops)
		case app.DestroyEvent:
			return e.Err
		}
	}
}

func (ui *UI) Layout(ctx layout.Context) layout.Dimensions {

	const (
		itemWidth  = 480 // 图片宽度
		itemHeight = 520 // 图片高度 + 文字区域高度
		minSpacing = 16  // 最小间距
	)

	// 计算每行可以放置的元素个数和间距
	availableWidth := ctx.Constraints.Max.X - ctx.Dp(unit.Dp(32)) // 减去左右边距
	itemWidthDp := ctx.Dp(unit.Dp(itemWidth))

	// 计算最大可能的列数（考虑最小间距）
	maxColumns := (availableWidth + ctx.Dp(unit.Dp(minSpacing))) / (itemWidthDp + ctx.Dp(unit.Dp(minSpacing)))
	columns := max(1, maxColumns)

	// 计算实际间距（包括两边的边距）
	totalGaps := columns + 1 // 元素之间的间隙数 + 两边的间隙
	spacing := max(ctx.Dp(unit.Dp(minSpacing)), (availableWidth-columns*itemWidthDp)/totalGaps)

	if ui.searchButton.Clicked(ctx) {
		fmt.Println("Search button clicked:", ui.search.Text())
	}

	wrapper := layout.Inset{
		Top:    dp24,
		Left:   dp16,
		Right:  dp16,
		Bottom: dp16,
	}
	return wrapper.Layout(ctx, func(ctx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(ctx,
			layout.Rigid(func(ctx layout.Context) layout.Dimensions {
				return layout.Center.Layout(ctx, func(ctx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(ctx,
						layout.Rigid(func(ctx layout.Context) layout.Dimensions {
							border := widget.Border{
								Color:        color.NRGBA{R: 0, G: 0, B: 0, A: 255},
								Width:        dp1,
								CornerRadius: dp8,
							}
							return border.Layout(ctx, func(ctx layout.Context) layout.Dimensions {
								ctx.Constraints.Min.X = ctx.Dp(unit.Dp(480))
								return layout.UniformInset(dp8).Layout(ctx, func(ctx layout.Context) layout.Dimensions {
									return material.Editor(ui.Theme, &ui.search, "请输入搜索内容...").Layout(ctx)
								})
							})
						}),
						layout.Rigid(layout.Spacer{Width: dp24}.Layout),
						layout.Rigid(func(ctx layout.Context) layout.Dimensions {
							return material.Button(ui.Theme, &ui.searchButton, "搜索影片").Layout(ctx)
						}),
					)
				})
			}),
			layout.Rigid(func(ctx layout.Context) layout.Dimensions {
				return layout.Spacer{Height: dp24}.Layout(ctx)
			}),
			layout.Flexed(1, func(ctx layout.Context) layout.Dimensions {
				// 设置最大宽度约束
				return component.Grid(ui.Theme, &ui.grid).Layout(ctx, 8, 1, // 使用单列
					func(axis layout.Axis, index, constraint int) int {
						if axis == layout.Horizontal {
							return constraint
						}
						return ctx.Dp(unit.Dp(itemHeight))
					},
					func(ctx layout.Context, row, _ int) layout.Dimensions {
						// 在每行内部实现水平居中的网格
						return layout.Flex{Axis: layout.Horizontal}.Layout(ctx,
							layout.Rigid(func(ctx layout.Context) layout.Dimensions {
								// 内部的水平网格
								return layout.Flex{Axis: layout.Horizontal}.Layout(ctx,
									ui.createGridItems(row, columns, itemWidthDp, spacing)...,
								)
							}),
						)
					},
				)
			}),
		)
	})
}

func (ui *UI) createGridItems(row, columns, itemWidthDp, spacing int) []layout.FlexChild {
	items := make([]layout.FlexChild, columns)
	for col := 0; col < columns; col++ {
		items[col] = layout.Rigid(func(ctx layout.Context) layout.Dimensions {
			// 第一列添加左边距，其他列添加右边距
			inset := layout.Inset{}
			if col == 0 {
				inset.Left = unit.Dp(float32(spacing))
			}
			inset.Right = unit.Dp(float32(spacing))

			return inset.Layout(ctx, func(ctx layout.Context) layout.Dimensions {
				return layout.Stack{Alignment: layout.N}.Layout(ctx,
					layout.Stacked(func(ctx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Vertical}.Layout(ctx,
							// 图片区域
							layout.Rigid(func(ctx layout.Context) layout.Dimensions {
								rect := clip.Rect{
									Max: image.Point{
										X: itemWidthDp,
										Y: itemWidthDp,
									},
								}.Op()

								paint.FillShape(ctx.Ops, color.NRGBA{
									R: uint8(255 / 8 * row),
									G: uint8(255 / columns * col),
									B: uint8(255 * row * col / (8 * columns)),
									A: 255,
								}, rect)

								return layout.Dimensions{
									Size: image.Point{
										X: itemWidthDp,
										Y: itemWidthDp,
									},
								}
							}),
							// 文字区域
							layout.Rigid(func(ctx layout.Context) layout.Dimensions {
								return layout.UniformInset(dp8).Layout(ctx,
									func(ctx layout.Context) layout.Dimensions {
										label := material.Label(ui.Theme, unit.Sp(14), fmt.Sprintf("Item %d-%d", row, col))
										return label.Layout(ctx)
									},
								)
							}),
						)
					}),
				)
			})
		})
	}
	return items
}
