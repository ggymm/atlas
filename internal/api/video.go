package api

import (
	"atlas/pkg/data/model"
	"atlas/pkg/data/service"
	"net/http"
)

type VideoApi struct {
	Api
}

type VideoResp struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Size      int64  `json:"size"`
	Path      string `json:"path"`
	Star      int64  `json:"star"`
	Tags      string `json:"tags"`
	Cover     string `json:"cover"`
	Duration  int64  `json:"duration"`
	UpdatedAt int64  `json:"updated_at"`
}

type VideoPageReq struct {
	*service.Page
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

	total, records, err := service.SelectVideos(req.Page, nil)
	if err != nil {
		internalServerError(w)
		return
	}
	videos := make([]*VideoResp, len(records))
	for i, v := range records {
		videos[i] = &VideoResp{
			Id:        v.Id,
			Name:      v.Name,
			Size:      v.Size,
			Path:      v.Path,
			Star:      v.Star,
			Tags:      v.Tags,
			Duration:  v.Duration,
			UpdatedAt: v.UpdatedAt,
		}
	}
	h.ok(w, VideoPageResp{Total: total, Records: videos})
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
	_, _ = w.Write(v.Cover)
}
