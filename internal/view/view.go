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
	theme *material.Theme

	spacing    unit.Dp
	itemWidth  unit.Dp
	itemHeight unit.Dp

	editor    widget.Editor
	searchBtn widget.Clickable
	drawerBtn widget.Clickable

	modalLayer  *component.ModalLayer
	modalDrawer *component.ModalNavDrawer

	currPage      int
	totalPage     int
	searchPrevBtn widget.Clickable
	searchNextBtn widget.Clickable

	searchResultGrid component.GridState
}

func NewUI() *UI {
	ui := new(UI)
	ui.theme = material.NewTheme()

	ui.editor = widget.Editor{
		SingleLine: true,
	}

	ui.modalLayer = component.NewModal()

	nav := component.NewNav("媒体库", "")
	ui.modalDrawer = component.ModalNavFrom(&nav, ui.modalLayer)
	ui.modalDrawer.AddNavItem(component.NavItem{
		Name: "测试",
	})
	ui.modalDrawer.AddNavItem(component.NavItem{
		Name: "测试2",
	})

	ui.spacing = unit.Dp(20)
	ui.itemWidth = unit.Dp(480)
	ui.itemHeight = unit.Dp(520)

	ui.currPage = 1
	ui.totalPage = 8
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
	for ui.drawerBtn.Clicked(ctx) {
		ui.modalDrawer.ToggleVisibility(ctx.Now)
	}

	if ui.searchBtn.Clicked(ctx) {
		fmt.Println("Search button clicked:", ui.editor.Text())
	}

	if ui.modalDrawer.NavDestinationChanged() {
		fmt.Println("Nav destination changed:", ui.modalDrawer.CurrentNavDestination())
	}

	// 处理分页按钮点击
	if ui.searchPrevBtn.Clicked(ctx) && ui.currPage > 1 {
		ui.currPage--
	}
	if ui.searchNextBtn.Clicked(ctx) && ui.currPage < ui.totalPage {
		ui.currPage++
	}

	layout.UniformInset(dp24).Layout(ctx, func(ctx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(ctx,
			// 搜索组件
			layout.Rigid(func(ctx layout.Context) layout.Dimensions {
				return layout.Center.Layout(ctx, func(ctx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(ctx,
						// 输入框
						layout.Rigid(func(ctx layout.Context) layout.Dimensions {
							border := widget.Border{
								Color:        color.NRGBA{R: 0, G: 0, B: 0, A: 255},
								Width:        dp1,
								CornerRadius: dp8,
							}
							return border.Layout(ctx, func(ctx layout.Context) layout.Dimensions {
								ctx.Constraints.Min.X = ctx.Dp(unit.Dp(480))
								ctx.Constraints.Max.X = ctx.Dp(unit.Dp(480))
								return layout.UniformInset(dp8).Layout(ctx, func(ctx layout.Context) layout.Dimensions {
									return material.Editor(ui.theme, &ui.editor, "请输入搜索内容...").Layout(ctx)
								})
							})
						}),
						layout.Rigid(layout.Spacer{Width: dp24}.Layout),
						// 搜索按钮
						layout.Rigid(func(ctx layout.Context) layout.Dimensions {
							return material.Button(ui.theme, &ui.searchBtn, "搜索影片").Layout(ctx)
						}),
						layout.Rigid(layout.Spacer{Width: dp24}.Layout),
						// 菜单按钮
						layout.Rigid(func(ctx layout.Context) layout.Dimensions {
							return material.Button(ui.theme, &ui.drawerBtn, "选择媒体库").Layout(ctx)
						}),
					)
				})
			}),

			// 搜索结果组件
			layout.Rigid(layout.Spacer{Height: dp24}.Layout),
			layout.Flexed(1, func(ctx layout.Context) layout.Dimensions {
				// grid 布局
				return component.Grid(ui.theme, &ui.searchResultGrid).Layout(ctx, 8, 1,
					func(axis layout.Axis, index, constraint int) int {
						if axis == layout.Horizontal {
							return constraint
						}
						return ctx.Dp(ui.itemHeight)
					},
					func(ctx layout.Context, row, _ int) layout.Dimensions {
						width := ctx.Constraints.Max.X
						spaceDp := ctx.Dp(ui.spacing)
						itemWithDp := ctx.Dp(ui.itemWidth)

						columns := max(1, (width+spaceDp)/(itemWithDp+spaceDp))         // 计算列数
						spacing := max(spaceDp, (width-columns*itemWithDp)/(columns+1)) // 计算间距

						return layout.Center.Layout(ctx, func(ctx layout.Context) layout.Dimensions {
							return layout.Flex{Axis: layout.Horizontal}.Layout(ctx,
								layout.Rigid(func(ctx layout.Context) layout.Dimensions {
									return layout.Flex{Axis: layout.Horizontal}.Layout(ctx,
										ui.createGridItems(row, columns, spacing)...,
									)
								}),
							)
						})
					},
				)
			}),

			// 搜索结果分页组件
			layout.Rigid(layout.Spacer{Height: dp16}.Layout),
			layout.Rigid(func(ctx layout.Context) layout.Dimensions {
				return layout.Center.Layout(ctx, func(ctx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(ctx,
						// 上一页按钮
						layout.Rigid(func(ctx layout.Context) layout.Dimensions {
							btn := new(widget.Clickable)
							if ui.currPage > 1 {
								btn = &ui.searchPrevBtn
							}
							return material.Button(ui.theme, btn, "上一页").Layout(ctx)
						}),
						// 页码信息
						layout.Rigid(func(ctx layout.Context) layout.Dimensions {
							return layout.Inset{
								Left:  dp24,
								Right: dp24,
							}.Layout(ctx, func(ctx layout.Context) layout.Dimensions {
								label := material.Label(ui.theme, unit.Sp(16),
									fmt.Sprintf("%d / %d", ui.currPage, ui.totalPage))
								return label.Layout(ctx)
							})
						}),
						// 下一页按钮
						layout.Rigid(func(ctx layout.Context) layout.Dimensions {
							btn := new(widget.Clickable)
							if ui.currPage < ui.totalPage {
								btn = &ui.searchNextBtn
							}
							return material.Button(ui.theme, btn, "下一页").Layout(ctx)
						}),
					)
				})
			}),
		)
	})

	ui.modalLayer.Layout(ctx, ui.theme)
	return layout.Dimensions{Size: ctx.Constraints.Max}
}

func (ui *UI) createGridItems(row, columns, spacing int) []layout.FlexChild {
	items := make([]layout.FlexChild, columns)
	for col := 0; col < columns; col++ {
		items[col] = layout.Rigid(func(ctx layout.Context) layout.Dimensions {
			inset := layout.Inset{}
			if col != 0 {
				inset.Left = unit.Dp(float32(spacing))
			}
			return inset.Layout(ctx, func(ctx layout.Context) layout.Dimensions {
				return layout.Stack{Alignment: layout.N}.Layout(ctx,
					layout.Stacked(func(ctx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Vertical}.Layout(ctx,
							// 图片区域
							layout.Rigid(func(ctx layout.Context) layout.Dimensions {
								rect := clip.Rect{
									Max: image.Point{
										X: ctx.Dp(ui.itemWidth),
										Y: ctx.Dp(ui.itemWidth),
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
										X: ctx.Dp(ui.itemWidth),
										Y: ctx.Dp(ui.itemWidth),
									},
								}
							}),
							// 文字区域
							layout.Rigid(func(ctx layout.Context) layout.Dimensions {
								return layout.UniformInset(dp8).Layout(ctx,
									func(ctx layout.Context) layout.Dimensions {
										label := material.Label(ui.theme, unit.Sp(14), fmt.Sprintf("Item %d-%d", row, col))
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
