package dto

type TrxResponse struct {
	ID          int64               `json:"id"`
	HargaTotal  int                 `json:"harga_total"`
	KodeInvoice string              `json:"kode_invoice"`
	MethodBayar string              `json:"method_bayar"`
	AlamatKirim AlamatResponse      `json:"alamat_kirim"`
	DetailTrx   []DetailTrxResponse `json:"detail_trx"`
}

type AlamatResponse struct {
	ID           int64  `json:"id"`
	JudulAlamat  string `json:"judul_alamat"`
	NamaPenerima string `json:"nama_penerima"`
	NoTelp       string `json:"no_telp"`
	DetailAlamat string `json:"detail_alamat"`
}

type DetailTrxResponse struct {
	Product    ProdukResponse `json:"product"`
	Toko       TokoResponse   `json:"toko"`
	Kuantitas  int            `json:"kuantitas"`
	HargaTotal int            `json:"harga_total"`
}

type ProdukResponseTRX struct {
	ID            int64            `json:"id"`
	NamaProduk    string           `json:"nama_produk"`
	Slug          string           `json:"slug"`
	HargaReseller int              `json:"harga_reseler"`
	HargaKonsumen int              `json:"harga_konsumen"`
	Deskripsi     string           `json:"deskripsi"`
	Toko          TokoResponse     `json:"toko"`
	Category      CategoryResponse `json:"category"`
	Photos        []PhotoResponse  `json:"photos"`
}

type TokoResponseTRX struct {
	ID       *int64  `json:"id"`
	NamaToko *string `json:"nama_toko"`
	UrlFoto  *string `json:"url_foto"`
}

type CategoryResponseTRX struct {
	ID           int64  `json:"id"`
	NamaCategory string `json:"nama_category"`
}

type PhotoResponseTRX struct {
	ID        int64  `json:"id"`
	ProductID int64  `json:"product_id"`
	Url       string `json:"url"`
}

type TrxListPaginated struct {
	Data  []TrxResponse `json:"data"`
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
}

type CreateTrxRequest struct {
	MethodBayar string                   `json:"method_bayar"`
	AlamatKirim int64                    `json:"alamat_kirim"`
	DetailTrx   []CreateDetailTrxRequest `json:"detail_trx"`
}

type CreateDetailTrxRequest struct {
	ProductID int64 `json:"product_id"`
	Kuantitas int   `json:"kuantitas"`
}