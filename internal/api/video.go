package api

import (
	"net/http"

	"atlas/pkg/data"
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

	Tags   string `json:"tags"`
	Order  string `json:"order"`
	Search string `json:"search"`
}

type VideoPageResp struct {
	Total   int64        `json:"total"`
	Records []*VideoResp `json:"records"`
}

func query(req *VideoPageReq) (*VideoPageResp, error) {
	var (
		size    = 20
		offset  = 1
		records []*model.Video
	)
	if req.PageReq != nil {
		size = req.PageReq.GetSize()
		offset = req.PageReq.GetOffset()
	}
	resp := new(VideoPageResp)

	// 查询总数
	err := data.DB.Model(&model.Video{}).Count(&resp.Total).Error
	if err != nil {
		return nil, err
	}

	// 查询列表
	err = data.DB.Limit(size).Offset(offset).Find(&records).Error
	if err != nil {
		return nil, err
	}

	// 视频列表
	resp.Records = make([]*VideoResp, len(records))
	for i, v := range records {
		resp.Records[i] = &VideoResp{
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
	return resp, nil
}

func search(req *VideoPageReq) (*VideoPageResp, error) {
	var (
		size    = 20
		offset  = 1
		records []*model.Video
	)
	if req.PageReq != nil {
		size = req.PageReq.GetSize()
		offset = req.PageReq.GetOffset()
	}
	resp := new(VideoPageResp)

	args := make([]any, 0)
	where := buildQuery(parseExpr(req.Search), &args)

	// 查询总数
	err := data.DB.Model(&model.Video{}).Where(where, args).Count(&resp.Total).Error
	if err != nil {
		return nil, err
	}

	// 查询列表
	err = data.DB.Limit(size).Offset(offset).Where(where, args).Find(&records).Error
	if err != nil {
		return nil, err
	}

	// 视频列表
	resp.Records = make([]*VideoResp, len(records))
	for i, v := range records {
		resp.Records[i] = &VideoResp{
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
	return resp, nil
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

	// 执行查询
	if len(req.Search) == 0 {
		resp, err = query(req)
	} else {
		resp, err = search(req)
	}
	if err != nil {
		internalServerError(w)
		return
	}
	h.ok(w, resp)
}
