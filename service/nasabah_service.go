package service

import (
	"corelyst-api/model"
	"corelyst-api/repository"
	"errors"
)

func GetAllNasabah() ([]model.Nasabah, error) {
	return repository.GetAllNasabah()
}

func AddNasabah(nasabah model.Nasabah) error {
	if nasabah.Nama == "" || nasabah.Alamat == "" || nasabah.Marketing == "" {
		return errors.New("nama, alamat & marketing wajib di isi")
	}

	return repository.AddNasabah(nasabah)
}
