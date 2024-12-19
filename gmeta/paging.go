package gmeta

type Paging struct {
	Limit  int   `json:"limit" validate:"max=1000"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total,omitempty"`

	BeforeValue string `json:"before_value,omitempty"`
	AfterValue  string `json:"after_value,omitempty"`
	NextValue   string `json:"next_value,omitempty"`
}

func (p *Paging) SetPage(page int) {
	p.Offset = page * p.Limit
}
