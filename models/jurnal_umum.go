package models

type JurnalUmum struct {
	ID              int     `json:"id"`
	Tanggal         string  `json:"tanggal"`
	NoRekening      string  `json:"no_rekening"`
	IDMasterAccount int     `json:"id_master_account"`
	Tipe            string  `json:"tipe"`
	Nominal         float64 `json:"nominal"`
	Keterangan      string  `json:"keterangan"`
	IDKaryawan      *int    `json:"id_karyawan"`
	InputBy         string  `json:"input_by"`
}
