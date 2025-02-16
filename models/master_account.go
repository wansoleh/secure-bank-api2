package models

// MasterAccount struct merepresentasikan data akun dalam sistem akuntansi bank
type MasterAccount struct {
	ID        int    `json:"id"`
	KodeAkun  string `json:"kode_akun"`
	NamaAkun  string `json:"nama_akun"`
	Tipe      string `json:"tipe"` // Aktiva, Pasiva, Ekuitas, Revenue, Beban, Dividen
}
