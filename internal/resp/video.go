package resp

import (
	"atlas/pkg/data/model"
)

type VideoPageResp struct {
	Total   int64          `json:"total"`
	Records []*model.Video `json:"records"`
}
