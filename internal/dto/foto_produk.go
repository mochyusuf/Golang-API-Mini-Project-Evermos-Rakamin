package dto

type FotoProdukResponse struct {
	ID        int64  `json:"id"`
	IdProduk  int64  `json:"id_produk"`
	Url       string `json:"url"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

type FotoProdukReq struct {
	IdProduk int64  `json:"id_produk"`
	Url      string `json:"url"`
}