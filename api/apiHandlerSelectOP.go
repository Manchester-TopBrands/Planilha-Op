package api

import (
	"encoding/json"
	"net/http"
	"op/models"	
)

type NewOP struct {
	Cod string `json:"cod,omitempty"`
}

func SelectOP(w http.ResponseWriter, r *http.Request) {
	NewOP := NewOP{}
	if err := json.NewDecoder(r.Body).Decode(&NewOP); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Status: "Bad Request",
			Error:  "",
			Data:   err.Error(),
		})
		return
	}
	
	data, err := connectionLinx.SelectOPDatabase(NewOP.Cod)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Status: "Bad Request",
			Error:  "",
			Data:   err.Error(),
		})
		return
	}

	if data == nil {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("access-control-expose-headers", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.ResponseExcell{
			Status: "Bad Request",
			Error:  "",
			Data:   "OP Não encontrada",
		})
		return 
	}
	excels := gerarExcel(data, NewOP.Cod)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("access-control-expose-headers", "*")
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("IdEx", excels)
	json.NewEncoder(w).Encode(models.ResponseExcell{
		Status: "OK",
		Error:  "",
		Data:   data,
		Id:     excels,
	})
}
