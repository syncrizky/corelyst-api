package model

type Nasabah struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Marketing string `json:"marketing"`
}
