package video

import (
	"os"
	"testing"

	"atlas/pkg/app"
)

func Test_Parse(t *testing.T) {
	app.Init()

	out, err := Parse("D:/temp/input.mp4")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", out)
}

func Test_Thumbnail(t *testing.T) {
	app.Init()

	output, err := Thumbnail("D:/temp/input.mp4")
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile("D:/temp/output.webp", output, 0666)
	if err != nil {
		t.Fatal(err)
	}
}
