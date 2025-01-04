package service

const (
	timeFormat = "2006-01-02 15:04:05"
)

type Get struct {
	Id int64 `json:"id" validate:"required"`
}

type Delete struct {
	Id int64 `json:"id" validate:"required"`
}

type Page struct {
	Page int `query:"page"`
	Size int `query:"size"`
}

func (p *Page) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Page) GetSize() int {
	if p.Size <= 0 {
		p.Size = 10
	}
	return p.Size
}

func (p *Page) GetOffset() int {
	return (p.GetPage() - 1) * p.GetSize()
}

type Option struct {
	Value int64  `json:"value"`
	Label string `json:"label"`
}

type PageList struct {
	List  any   `json:"list"`
	Total int64 `json:"total"`
}

type TreeNode[T any] struct {
	Key      int64  `json:"key"`
	Label    string `json:"label"`
	Children []*T   `json:"children"`
}
