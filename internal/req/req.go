package req

type BasePage struct {
	Page int `query:"page"`
	Size int `query:"size"`
}

func (p *BasePage) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *BasePage) GetSize() int {
	if p.Size <= 0 {
		p.Size = 10
	}
	return p.Size
}

func (p *BasePage) GetOffset() int {
	return (p.GetPage() - 1) * p.GetSize()
}

type BaseGet struct {
	Id int64 `json:"id" validate:"required"`
}

type BaseDelete struct {
	Id int64 `json:"id" validate:"required"`
}
