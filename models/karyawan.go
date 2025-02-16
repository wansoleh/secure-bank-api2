package models

// Karyawan struct merepresentasikan data karyawan dalam database
type Karyawan struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Nama      string `json:"nama"`
	NIP       string `json:"nip"`
	Posisi    string `json:"posisi"`
	CreatedAt string `json:"created_at"`
}
