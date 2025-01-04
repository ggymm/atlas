package view

import (
	"os"

	"gioui.org/app"
	"gioui.org/unit"
	"gioui.org/widget"
)

var (
	dp1  = unit.Dp(1)
	dp8  = unit.Dp(8)
	dp16 = unit.Dp(16)
	dp24 = unit.Dp(24)
	dp40 = unit.Dp(40)

	sp12 = unit.Sp(12)
	sp14 = unit.Sp(14)
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

	nameLayout  LabelLayout
	coverLayout ImageLayout

	videoClickable widget.Clickable
}

func Show() {
	ui := NewUI()
	go func() {
		w := new(app.Window)
		w.Option(
			app.Title("Atlas"),
			app.Size(unit.Dp(1200), unit.Dp(800)),    // 设置默认大小
			app.MinSize(unit.Dp(1200), unit.Dp(800)), // 设置最小大小
		)
		if err := ui.Run(w); err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	}()
	app.Main()
}
