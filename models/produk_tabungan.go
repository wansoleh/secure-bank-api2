package models

// ProdukTabungan struct merepresentasikan data produk tabungan dalam database
type ProdukTabungan struct {
	ID               int     `json:"id"`
	NamaProduk       string  `json:"nama_produk"`
	Deskripsi        string  `json:"deskripsi"`
	Bunga            float64 `json:"bunga"`              // Persentase bunga tahunan
	BiayaAdministrasi float64 `json:"biaya_administrasi"` // Biaya admin bulanan
	MinSaldo         float64 `json:"min_saldo"`          // Saldo minimum untuk produk tabungan
	CreatedAt        string  `json:"created_at"`
}
