
package models

// LogAktivitas struct merepresentasikan tabel log_aktivitas
type LogAktivitas struct {
	ID         int    `json:"id"`
	IDJurnal   int    `json:"id_jurnal"`
	IDKaryawan *int   `json:"id_karyawan"`
	InputBy    string `json:"input_by"`
	Aksi       string `json:"aksi"`
	Timestamp  string `json:"timestamp"`
	Keterangan string `json:"keterangan"`
}
