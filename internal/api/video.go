package api

import (
	"atlas/pkg/app"
	"atlas/pkg/data"
	"atlas/pkg/data/model"
	"atlas/pkg/data/service"
	"atlas/pkg/utils"
	"log/slog"
	"net/http"
	"path/filepath"
	"slices"
)

type VideoApi struct {
	Api
}

type VideoResp struct {
	Id        string `gorm:"column:id" json:"id"`
	Path      string `gorm:"column:path" json:"path"`
	Size      int64  `gorm:"column:size" json:"size"`
	Tags      string `gorm:"column:tags" json:"tags"`
	Title     string `gorm:"column:title" json:"title"`
	Stars     int64  `gorm:"column:stars" json:"stars"`
	Format    string `gorm:"column:format" json:"cover"`
	Duration  int64  `gorm:"column:duration" json:"duration"`
	UpdatedAt int64  `gorm:"column:updated_at" json:"updated_at"`
}

func (*VideoResp) TableName() string {
	return "video"
}

type VideoPageReq struct {
	Page   int    `json:"page"`
	Size   int    `json:"size"`
	Tags   string `json:"tags"`
	Path   string `json:"path"`
	Order  string `json:"order"`
	Search string `json:"search"`
}

type VideoPageResp struct {
	Total   int64        `json:"total"`
	Records []*VideoResp `json:"records"`
}

type VideoInfoResp struct {
	Root          string `json:"root"`
	Total         int64  `json:"total"`
	TotalSize     string `json:"totalSize"`
	TotalDuration string `json:"totalDuration"`
}

type VideoUpdateReq struct {
	Id    string `json:"id"`
	Stars int64  `json:"stars"`
}

func (h *VideoApi) Cover(w http.ResponseWriter, r *http.Request) {
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

func (h *VideoApi) QueryInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}

	ret := struct {
		Total         int64 `gorm:"column:total"`
		TotalSize     int64 `gorm:"column:total_size"`
		TotalDuration int64 `gorm:"column:total_duration"`
	}{}
	fields := "count(id) as total, sum(size) as total_size, sum(duration) as total_duration"
	err := data.DB.Model(&model.Video{}).Select(fields).First(&ret).Error
	if err != nil {
		internalServerError(w)
		return
	}

	info := new(VideoInfoResp)
	info.Root = app.Root
	info.Total = ret.Total
	info.TotalSize = utils.FormatSize(ret.TotalSize)
	info.TotalDuration = utils.FormatDuration(ret.TotalDuration)
	h.ok(w, info)
}

func (h *VideoApi) QueryPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}
	req := new(VideoPageReq)
	resp := new(VideoPageResp)

	// 解析请求
	err := ParseJSON(r, &req)
	if err != nil {
		badRequest(w)
		return
	}

	var (
		limit  = 20
		offset = 0
	)
	if req.Page != 0 && req.Size != 0 {
		limit = req.Size
		offset = (req.Page - 1) * req.Size
	}
	total := data.DB.Model(&model.Video{})
	records := data.DB.Limit(limit).Offset(offset)

	if len(req.Search) != 0 {
		args := make([]any, 0)
		query := buildQuery(parseExpr(req.Search), &args)

		slog.Info("search expr parsed", slog.String("query", query))

		total = total.Where(query, args...)
		records = records.Where(query, args...)
	}

	// 查询总数
	err = total.Count(&resp.Total).Error
	if err != nil {
		internalServerError(w)
		return
	}

	// 查询列表
	err = records.Find(&resp.Records).Error
	if err != nil {
		internalServerError(w)
		return
	}
	h.ok(w, resp)
}

func (h *VideoApi) QueryPaths(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}

	ps := make([]string, 0)
	err := data.DB.Model(&model.Video{}).Select("path").Pluck("path", &ps).Error
	if err != nil {
		internalServerError(w)
		return
	}

	paths := make([]string, 0)
	for _, p := range ps {
		p = filepath.Dir(p)
		p = filepath.Base(p)
		if !slices.Contains(paths, p) {
			paths = append(paths, p)
		}
	}
	h.ok(w, paths)
}

func (h *VideoApi) UpdateStars(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}
	req := new(VideoUpdateReq)

	// 解析请求
	err := ParseJSON(r, &req)
	if err != nil {
		badRequest(w)
		return
	}

	// 更新评分
	video := &model.Video{
		Id: req.Id,
	}
	update := map[string]any{
		"stars": req.Stars,
	}
	err = data.DB.Model(video).Updates(update).Error
	if err != nil {
		internalServerError(w)
		return
	}
	h.ok(w, true)
}
