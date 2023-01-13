package products

type PageInfo struct {
	TotalData int    `json:"total_data"`
	Page      int    `json:"page"`
	LastPage  int    `json:"last_page"`
	EmptyPage string `json:"empty_page"`
}
