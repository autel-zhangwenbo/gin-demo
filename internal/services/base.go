package services


type PageParam struct {
	Page int `form:"page" binding:"required,gte=1"`
	Size int `form:"size" binding:"required,lte=1000"`
}

type PageResponse struct {
	Total int64 `json:"total"`
	Data  any   `json:"data,omitempty"`
}
