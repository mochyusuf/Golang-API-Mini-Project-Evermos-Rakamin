package dto

import "time"

type TokoResponse struct {
	ID        int64     `json:"id"`
	IdUser    int64     `json:"id_user"`
	NamaToko  string    `json:"nama_toko"`
	UrlFoto   string    `json:"url_foto"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateTokoRequest struct {
	IdUser    int64  `json:"id_user"`

}

type UpdateTokoRequest struct {
	ID       int64   `json:"id"`
	NamaToko *string `json:"nama_toko"`
	UrlFoto  *string `json:"url_foto"`
}
