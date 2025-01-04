package task

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/ggymm/gopkg/xxhash"

	"atlas/pkg/log"
)

func check(p string, h *xxhash.Digest) string {
	h.Reset()

	// 检查格式
	f, err := os.Open(p)
	if err != nil {
		log.Error(err).Str("file", p).Msg("file open error")
		return ""
	}
	defer func() {
		_ = f.Close()
	}()

	buf := make([]byte, 512)
	_, err = f.Read(buf)
	if err != nil {
		log.Error(err).Str("file", p).Msg("file read error")
		return ""
	}

	mime := http.DetectContentType(buf)
	if !strings.HasPrefix(mime, "video/") {
		log.Error(err).Str("file", p).Msg("file format error")
		return ""
	}

	// 计算文件 hash
	_, err = io.Copy(h, f)
	if err != nil {
		log.Error(err).Str("file", p).Msg("file xxhash error")
		return ""
	}
	return h.SumHex()
}
