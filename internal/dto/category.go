package dto

type CategoryResponse struct {
	ID           int64  `json:"id"`
	NamaCategory string `json:"nama_category"`
}

type CreateCategoryRequest struct {
	NamaCategory string `json:"nama_category"`
}

type UpdateCategoryRequest struct {
	NamaCategory *string `json:"nama_category"`
}