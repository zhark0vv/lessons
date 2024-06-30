package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type service interface {
	ProcessData(ctx context.Context, id int) (string, error)
}

type API struct {
	service service
}

func New(svc service) *API {
	return &API{service: svc}
}

func (api *API) GetDataHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	data, err := api.service.ProcessData(ctx, id)
	if err != nil {
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"data": data})
}

func (api *API) Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/data/{id}", api.GetDataHandler).Methods("GET")
	return r
}
