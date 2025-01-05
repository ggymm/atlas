package task

import (
	"atlas/pkg/app"
)

type Scanner struct {
	root string
}

func NewScanner() *Scanner {
	return &Scanner{
		root: app.Root,
	}
}
