package task_test

import (
	"net/http"
	"os"
	"testing"
)

func Test_Mime(t *testing.T) {
	f, err := os.Open("D:/temp/input.mp4")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = f.Close()
	}()

	buf := make([]byte, 512)
	_, err = f.Read(buf)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("mime: %s", http.DetectContentType(buf))
}
