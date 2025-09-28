package dto

type ProdukResponse struct {
	ID            int64            `json:"id"`
	NamaProduk    string           `json:"nama_produk"`
	Slug          string           `json:"slug"`
	HargaReseller int              `json:"harga_reseler"`
	HargaKonsumen int              `json:"harga_konsumen"`
	Stok          int              `json:"stok"`
	Deskripsi     string           `json:"deskripsi"`
	Toko          TokoResponse     `json:"toko"`
	Category      CategoryResponse `json:"category"`
	Photos        []PhotoResponse  `json:"photos"`
}

type PhotoResponse struct {
	ID        int64  `json:"id"`
	ProductID int64  `json:"product_id"`
	Url       string `json:"url"`
}

type CreateProdukRequest struct {
	NamaProduk    string `json:"nama_produk"`
	Slug          string `json:"slug"`
	HargaReseller int    `json:"harga_reseller"`
	HargaKonsumen int    `json:"harga_konsumen"`
	Stok          int    `json:"stok"`
	Deskripsi     string `json:"deskripsi"`
	IdToko        int64  `json:"id_toko"`
	IdCategory    int64  `json:"id_category"`
}

type UpdateProdukRequest struct {
	NamaProduk    *string `json:"nama_produk"`
	Slug          *string `json:"slug"`
	HargaReseller *int    `json:"harga_reseller"`
	HargaKonsumen *int    `json:"harga_konsumen"`
	Stok          *int    `json:"stok"`
	Deskripsi     *string `json:"deskripsi"`
	IdCategory    *int64  `json:"id_category"`
}

type ProdukFilterParams struct {
	NamaProduk string
	Limit      int
	Page       int
	CategoryID int64
	TokoID     int64
	MinHarga   int
	MaxHarga   int
}

type ProdukListPaginated struct {
	Data  []ProdukResponse `json:"data"`
	Page  int              `json:"page"`
	Limit int              `json:"limit"`
}