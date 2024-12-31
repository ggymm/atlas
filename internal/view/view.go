package view

import (
	_ "golang.org/x/image/webp"

	"bytes"
	"fmt"
	"image"
	"image/color"

	"gioui.org/app"
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"atlas/pkg/data/service"
)

type UI struct {
	win   *app.Window
	theme *material.Theme

	editor    widget.Editor
	searchBtn widget.Clickable
	drawerBtn widget.Clickable

	page    int
	total   int
	pageNum int

	width   unit.Dp
	height  unit.Dp
	spacing unit.Dp

	loading     bool
	records     []*Video
	recordsGrid component.GridState

	searchPrevBtn widget.Clickable
	searchNextBtn widget.Clickable

	layer  *component.ModalLayer
	drawer *component.ModalNavDrawer
}

func NewUI() *UI {
	ui := new(UI)
	ui.theme = material.NewTheme()

	ui.editor = widget.Editor{
		SingleLine: true,
	}

	ui.page = 1

	ui.width = unit.Dp(320)  // 1920 * 1080 => 320 * 180
	ui.height = unit.Dp(220) // 180 + 40
	ui.spacing = unit.Dp(20)

	ui.layer = component.NewModal()

	nav := component.NewNav("媒体库", "")
	ui.drawer = component.ModalNavFrom(&nav, ui.layer)
	ui.drawer.AddNavItem(component.NavItem{
		Name: "测试",
	})
	ui.drawer.AddNavItem(component.NavItem{
		Name: "测试2",
	})

	go ui.Load()
	return ui
}

func (ui *UI) Run(w *app.Window) error {
	ui.win = w
	ops := new(op.Ops)
	for {
		switch e := w.Event().(type) {
		case app.FrameEvent:
			ctx := app.NewContext(ops, e)

			if ui.drawerBtn.Clicked(ctx) {
				ui.drawer.ToggleVisibility(ctx.Now)
			}

			if ui.searchBtn.Clicked(ctx) {
				fmt.Println("Search button clicked:", ui.editor.Text())
			}
			if ui.searchPrevBtn.Clicked(ctx) && ui.page > 1 {
				ui.page--
				go ui.Load()
			}
			if ui.searchNextBtn.Clicked(ctx) && ui.page < ui.pageNum {
				ui.page++
				go ui.Load()
			}
			ui.Layout(ctx)
			e.Frame(ctx.Ops)
		case app.DestroyEvent:
			return e.Err
		}
	}
}

func (ui *UI) Load() {
	ui.loading = true
	size := 20
	data, err := service.FetchVideos(&service.VideoPageReq{
		Page: service.Page{
			Page: ui.page,
			Size: size, // 每页显示数量
		},
	})
	ui.loading = false
	if err != nil {
		fmt.Println("FetchVideos error:", err)
		return
	}

	videos := make([]*Video, len(data.Records))
	for i, r := range data.Records {
		// 原始数据
		v := &Video{
			Id:        r.Id,
			Name:      r.Name,
			Tags:      r.Tags,
			Path:      r.Path,
			Size:      r.Size,
			Cover:     r.Cover,
			Format:    r.Format,
			Duration:  r.Duration,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		}

		webpImg, _, _ := image.Decode(bytes.NewReader(r.Cover))
		if webpImg != nil {
			v.coverImage = webpImg
			v.coverImageOp = paint.NewImageOp(webpImg)
		}

		v.nameTip = component.DesktopTooltip(ui.theme, v.Name)
		v.nameTipArea = &component.TipArea{}
		videos[i] = v
	}

	ui.total = int(data.Total)
	ui.records = videos

	ui.pageNum = (ui.total + size - 1) / size // 计算总页数

	// 刷新界面
	if ui.win != nil {
		ui.win.Invalidate()
	}
}

func (ui *UI) Layout(ctx layout.Context) layout.Dimensions {
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
				if ui.loading {
					// 显示加载动画
					return layout.Center.Layout(ctx, func(ctx layout.Context) layout.Dimensions {
						return material.Loader(ui.theme).Layout(ctx)
					})
				}

				width := ctx.Constraints.Max.X
				withDp := ctx.Dp(ui.width)
				spaceDp := ctx.Dp(ui.spacing)

				columns := max(1, (width+spaceDp)/(withDp+spaceDp))         // 计算列数
				spacing := max(spaceDp, (width-columns*withDp)/(columns+1)) // 计算间距

				rows := (len(ui.records) + columns - 1) / columns // 向上取整计算行数

				// grid 布局
				return component.Grid(ui.theme, &ui.recordsGrid).Layout(ctx, rows, 1,
					func(axis layout.Axis, index, constraint int) int {
						if axis == layout.Horizontal {
							return constraint
						}
						return ctx.Dp(ui.height)
					},
					func(ctx layout.Context, row, _ int) layout.Dimensions {
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
							if ui.page > 1 {
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
									fmt.Sprintf("%d / %d", ui.page, ui.pageNum))
								return label.Layout(ctx)
							})
						}),

						// 下一页按钮
						layout.Rigid(func(ctx layout.Context) layout.Dimensions {
							btn := new(widget.Clickable)
							if ui.page < ui.pageNum {
								btn = &ui.searchNextBtn
							}
							return material.Button(ui.theme, btn, "下一页").Layout(ctx)
						}),
					)
				})
			}),
		)
	})

	ui.layer.Layout(ctx, ui.theme)
	return layout.Dimensions{Size: ctx.Constraints.Max}
}

func (ui *UI) createGridItems(row, columns, spacing int) []layout.FlexChild {
	items := make([]layout.FlexChild, columns)
	for col := 0; col < columns; col++ {
		i := row*columns + col
		if i >= len(ui.records) {
			// 渲染空布局
			items[col] = layout.Rigid(func(ctx layout.Context) layout.Dimensions {
				return layout.Inset{Left: unit.Dp(float32(spacing))}.Layout(ctx,
					func(ctx layout.Context) layout.Dimensions {
						return layout.Dimensions{
							Size: image.Point{
								X: ctx.Dp(ui.width),
								Y: ctx.Dp(ui.height),
							},
						}
					},
				)
			})
			continue
		}

		// 渲染真实布局
		video := ui.records[i]
		items[col] = layout.Rigid(func(ctx layout.Context) layout.Dimensions {
			inset := layout.Inset{}
			if col != 0 {
				inset.Left = unit.Dp(float32(spacing))
			}
			// 处理点击事件
			if video.clickable.Clicked(ctx) {
				fmt.Printf("Opening video: %s\n", video.Path)
			}
			return video.clickable.Layout(ctx, func(ctx layout.Context) layout.Dimensions {
				return inset.Layout(ctx, func(ctx layout.Context) layout.Dimensions {
					return layout.Stack{Alignment: layout.N}.Layout(ctx,
						layout.Stacked(func(ctx layout.Context) layout.Dimensions {
							return layout.Flex{Axis: layout.Vertical}.Layout(ctx,
								// 图片区域
								layout.Rigid(func(ctx layout.Context) layout.Dimensions {
									width := ctx.Dp(ui.width)
									height := ctx.Dp(ui.height - dp40)
									if video.coverImage != nil {
										video.coverImageOp.Add(ctx.Ops)
										paint.PaintOp{}.Add(ctx.Ops)
									} else {
										rect := clip.Rect{Max: image.Point{X: width, Y: height}}.Op()
										paint.FillShape(ctx.Ops, color.NRGBA{R: 200, G: 200, B: 200, A: 255}, rect)
									}
									return layout.Dimensions{Size: image.Point{X: width, Y: height}}
								}),

								// 文字区域
								layout.Rigid(func(ctx layout.Context) layout.Dimensions {
									return layout.UniformInset(dp8).Layout(ctx,
										func(ctx layout.Context) layout.Dimensions {
											ctx.Constraints.Min.X = ctx.Dp(ui.width - dp16)
											ctx.Constraints.Max.X = ctx.Dp(ui.width - dp16)
											// ctx.Constraints.Min.Y = ctx.Dp(20)
											ctx.Constraints.Max.Y = ctx.Dp(ui.height - dp40)

											// tooltip
											tip := video.nameTip
											area := video.nameTipArea
											return area.Layout(ctx, tip, func(ctx layout.Context) layout.Dimensions {
												ctx.Constraints.Min.X = ctx.Dp(ui.width - dp16)
												ctx.Constraints.Max.X = ctx.Dp(ui.width - dp16)

												label := material.Label(ui.theme, sp12, video.Name)
												label.MaxLines = 1
												label.Alignment = text.Middle
												label.Font.Weight = font.Bold
												return label.Layout(ctx)
											})
										},
									)
								}),
							)
						}),
					)
				})
			})
		})
	}
	return items
}
