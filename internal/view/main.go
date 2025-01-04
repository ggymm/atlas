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
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"atlas/pkg/data/service"
	"atlas/pkg/log"
)

const (
	itemSpace  = unit.Dp(20)
	itemWidth  = unit.Dp(320) // 1920 * 1080 => 320 * 180
	itemHeight = unit.Dp(220) // 180 + 40
)

type UI struct {
	win   *app.Window
	theme *material.Theme

	layer  *component.ModalLayer
	drawer *component.ModalNavDrawer

	search    widget.Editor
	searchBtn widget.Clickable
	drawerBtn widget.Clickable

	page    int
	total   int
	pageNum int

	loading     bool
	records     []*Video
	recordsGrid component.GridState

	searchPrevBtn widget.Clickable
	searchNextBtn widget.Clickable
}

func NewUI() *UI {
	ui := new(UI)
	ui.theme = material.NewTheme()

	ui.search = widget.Editor{
		SingleLine: true,
	}

	ui.page = 1

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
				fmt.Println("Search button clicked:", ui.search.Text())
			}
			if ui.searchPrevBtn.Clicked(ctx) && ui.page > 1 {
				ui.page--
				ui.recordsGrid = component.GridState{}
				go ui.Load()
			}
			if ui.searchNextBtn.Clicked(ctx) && ui.page < ui.pageNum {
				ui.page++
				ui.recordsGrid = component.GridState{}
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
	size := 60
	data, err := service.SelectVideos(&service.VideoPageReq{
		Page: service.Page{
			Page: ui.page,
			Size: size, // 每页显示数量
		},
	})
	ui.loading = false
	if err != nil {
		log.Error(err).Msg("fetch data error")
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

		label := material.Label(ui.theme, sp14, v.Name)
		label.MaxLines = 1
		label.Alignment = text.Middle
		label.Font.Weight = font.Bold
		v.nameLayout = LabelLayout{
			LabelStyle: label,
			State:      &component.TipArea{},
			Tooltip:    component.DesktopTooltip(ui.theme, v.Name),
		}

		v.coverLayout = ImageLayout{}
		webpImg, _, _ := image.Decode(bytes.NewReader(r.Cover))
		if webpImg != nil {
			v.coverLayout.image = webpImg
			v.coverLayout.imageOp = paint.NewImageOp(webpImg)
		}
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
									return material.Editor(ui.theme, &ui.search, "请输入搜索内容...").Layout(ctx)
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
				spaceDp := ctx.Dp(itemSpace)
				widthDp := ctx.Dp(itemWidth)

				columns := max(1, (width+spaceDp)/(widthDp+spaceDp))         // 计算列数
				spacing := max(spaceDp, (width-columns*widthDp)/(columns+1)) // 计算间距

				rows := (len(ui.records) + columns - 1) / columns // 向上取整计算行数

				// grid 布局
				return component.Grid(ui.theme, &ui.recordsGrid).Layout(ctx, rows, 1,
					func(axis layout.Axis, index, constraint int) int {
						if axis == layout.Horizontal {
							return constraint
						}
						return ctx.Dp(itemHeight)
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
								label := material.Label(ui.theme, sp14, fmt.Sprintf("%d / %d", ui.page, ui.pageNum))
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
	return layout.Dimensions{}
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
								X: ctx.Dp(itemWidth),
								Y: ctx.Dp(itemHeight),
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
			if video.videoClickable.Clicked(ctx) {
				fmt.Printf("Opening video: %s\n", video.Path)
			}
			return video.videoClickable.Layout(ctx, func(ctx layout.Context) layout.Dimensions {
				return inset.Layout(ctx, func(ctx layout.Context) layout.Dimensions {
					return layout.Stack{Alignment: layout.N}.Layout(ctx,
						layout.Stacked(func(ctx layout.Context) layout.Dimensions {
							return layout.Flex{Axis: layout.Vertical}.Layout(ctx,
								// 图片区域
								layout.Rigid(func(ctx layout.Context) layout.Dimensions {
									width := ctx.Dp(itemWidth)
									height := ctx.Dp(itemHeight - dp40)
									return video.coverLayout.Layout(ctx, width, height)
								}),

								// 文字区域
								layout.Rigid(func(ctx layout.Context) layout.Dimensions {
									return layout.UniformInset(dp8).Layout(ctx,
										func(ctx layout.Context) layout.Dimensions {
											ctx.Constraints.Min.X = ctx.Dp(itemWidth - dp16)
											ctx.Constraints.Max.X = ctx.Dp(itemWidth - dp16)
											// ctx.Constraints.Min.Y = ctx.Dp(20)
											ctx.Constraints.Max.Y = ctx.Dp(itemHeight - dp40)

											return video.nameLayout.Layout(ctx)
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
