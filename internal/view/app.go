package view

import (
	"atlas/pkg/app"
	"os/exec"
)

type Webview struct {
	Bin  string
	View string
}

func NewWebview() *Webview {
	return &Webview{
		Bin:  app.Webview,
		View: app.View,
	}
}

func (w *Webview) Start() error {
	cmd := exec.Command(app.Webview, app.View)
	err := cmd.Start()
	if err != nil {
		return err
	}
	return cmd.Wait()
}
