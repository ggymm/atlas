package main

import (
	"os"

	"atlas/pkg/app"
	"atlas/pkg/movie"
)

func init() {
	app.Init()
}

func main() {
	bs, err := movie.Thumbnail("D:\\temp\\input.mp4")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("D:\\temp\\foo.jpg", bs, 0644)
	if err != nil {
		panic(err)
	}
}
