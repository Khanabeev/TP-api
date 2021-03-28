package apiserver

import (
	"net/http"

	"github.com/Khanabeev/TP-api/internal/app/service"
	"github.com/gorilla/mux"
)

type MasterHandler struct {
	Service service.MasterService
}

func (mh *MasterHandler) GetMasterById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	masterId := vars["master_id"]

	master, errApp := mh.Service.GetMasterById(masterId)
	if errApp != nil {
		writeResponse(w, errApp.Code, errApp.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, master)
}
