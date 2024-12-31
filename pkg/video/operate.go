package video

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"

	"atlas/pkg/app"
)

func Parse(path string) (*Info, error) {
	cmd := exec.Command(app.Ffprobe, "-hide_banner", "-v", "error",
		"-select_streams", "v:0", "-show_format", "-print_format", "json", path,
	)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// 解析 JSON
	i := new(Info)
	err = json.Unmarshal(out, &i)
	if err != nil {
		return nil, err
	}
	return i, err
}

func Thumbnail(path string) ([]byte, error) {
	buf := new(bytes.Buffer)
	cmd := exec.Command(app.Ffmpeg, "-hide_banner", "-v", "error",
		"-i", path, "-vf", "scale=320:180:force_original_aspect_ratio=decrease",
		"-c:v", "webp", "-preset", "picture", "-q:v", "80", "-frames:v", "1", "-f", "image2pipe", "-",
	)

	cmd.Stdout = buf
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
