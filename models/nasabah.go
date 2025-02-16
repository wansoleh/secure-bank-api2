package models

// Nasabah struct merepresentasikan data nasabah dalam database
type Nasabah struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	Nama        string  `json:"nama"`
	NIK         string  `json:"nik"`
	NoHP        string  `json:"no_hp"`
	NoRekening  string  `json:"no_rekening"`
	NoPIN       string  `json:"no_pin,omitempty"`  // PIN tidak akan ditampilkan di response JSON
	Saldo       float64 `json:"saldo"`
	CreatedAt   string  `json:"created_at"`
}
