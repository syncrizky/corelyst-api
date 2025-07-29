package repository

import (
	"corelyst-api/config"
	"corelyst-api/model"
)

func GetAllNasabah() ([]model.Nasabah, error) {
	rows, err := config.DB.Query("SELECT id, nama, alamat, marketing FROM nasabah")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allNasabah []model.Nasabah
	for rows.Next() {
		var ns model.Nasabah
		rows.Scan(&ns.ID, &ns.Nama, &ns.Alamat, &ns.Marketing)
		allNasabah = append(allNasabah, ns)
	}
	return allNasabah, nil
}

func AddNasabah(nasabah model.Nasabah) error {
	_, err := config.DB.Exec("INSERT INTO nasabah(nama, alamat, marketing) VALUES (?, ?, ?)", nasabah.Nama, nasabah.Alamat, nasabah.Marketing)
	return err
}
