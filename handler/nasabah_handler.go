package handler

import (
	"corelyst-api/model"
	"corelyst-api/service"
	"encoding/json"
	"net/http"
)

func GetAllNasabahHandler(w http.ResponseWriter, r *http.Request) {
	allNasabah, _ := service.GetAllNasabah()
	json.NewEncoder(w).Encode(allNasabah)
}

func AddNasabahHandler(w http.ResponseWriter, r *http.Request) {
	var nasabah model.Nasabah
	json.NewDecoder(r.Body).Decode(&nasabah)

	err := service.AddNasabah(nasabah)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("nasabah berhasil di daftarkan"))
}
