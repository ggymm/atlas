package video

import (
	"bytes"
	"os"
	"os/exec"

	"atlas/pkg/app"
)

func Thumbnail(input string) ([]byte, error) {
	buf := bytes.Buffer{}

	args := []string{
		"-hide_banner", "-v", "error", "-y",
		"-i", input,
		"-vf", "scale=480:480:force_original_aspect_ratio=decrease",
		"-c:v", "mjpeg", "-q:v", "5",
		"-frames:v", "1", "-f", "image2pipe", "-", "-f", "mjpeg",
	}
	cmd := exec.Command(app.Ffmpeg, args...)
	cmd.Stdout = &buf
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
