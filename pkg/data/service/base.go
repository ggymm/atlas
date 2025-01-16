package service

type PageReq struct {
	Page int
	Size int
}

func (p *PageReq) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *PageReq) GetSize() int {
	if p.Size <= 0 {
		p.Size = 10
	}
	return p.Size
}

func (p *PageReq) GetOffset() int {
	return (p.GetPage() - 1) * p.GetSize()
}

type PageResp[T any] struct {
	Total   int64
	Records []*T
}

type OptionResp struct {
	Value int64
	Label string
}

type TreeNodeResp[T any] struct {
	Key      int64
	Label    string
	Children []*T
}
