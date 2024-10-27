package internal

type Pagination struct {
	PageNo   int `form:"pageNo" json:"pageNo"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

func (m *Pagination) PageOffset() int {
	if m.PageNo <= 0 {
		m.PageNo = 1
	}
	if m.PageSize <= 0 {
		m.PageSize = 10
	}
	return (m.PageNo - 1) * m.PageSize
}

func (m *Pagination) PageLimit() int {
	if m.PageSize <= 0 {
		m.PageSize = 10
	}
	return m.PageSize
}

type PageResult struct {
	Data      interface{} `json:"data"`
	PageIndex int         `form:"page" json:"page"`
	PageSize  int         `form:"size" json:"size"`
	Total     int         `form:"total" json:"total"`
}
