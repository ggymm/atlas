package api

import (
	"net/http"

	"atlas/pkg/data/model"
	"atlas/pkg/data/service"
)

type VideoApi struct {
	Api
}

type VideoResp struct {
	Id        string `json:"id"`
	Path      string `json:"path"`
	Size      int64  `json:"size"`
	Star      int64  `json:"star"`
	Tags      string `json:"tags"`
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	Duration  int64  `json:"duration"`
	UpdatedAt int64  `json:"updated_at"`
}

type VideoPageReq struct {
	*service.PageReq
}

type VideoPageResp struct {
	Total   int64        `json:"total"`
	Records []*VideoResp `json:"records"`
}

func (h *VideoApi) GetPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}

	req := new(VideoPageReq)
	err := ParseJSON(r, &req)
	if err != nil {
		badRequest(w)
		return
	}

	resp, err := service.QueryVideos(req.PageReq)
	if err != nil {
		internalServerError(w)
		return
	}
	videos := make([]*VideoResp, len(resp.Records))
	for i, v := range resp.Records {
		videos[i] = &VideoResp{
			Id:        v.Id,
			Path:      v.Path,
			Size:      v.Size,
			Star:      v.Star,
			Tags:      v.Tags,
			Title:     v.Title,
			Duration:  v.Duration,
			UpdatedAt: v.UpdatedAt,
		}
	}
	h.ok(w, VideoPageResp{Total: resp.Total, Records: videos})
}

func (h *VideoApi) GetCover(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}

	v := new(model.Video)
	v.Id = r.PathValue("id")
	err := service.GetVideo(v)
	if err != nil {
		internalServerError(w)
		return
	}

	// 输出封面
	w.Header().Set("Content-Type", "image/webp")
	w.Header().Set("Cache-Control", "no-store") // 禁止缓存，减少浏览器内存占用
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	_, _ = w.Write(v.Cover)
}
