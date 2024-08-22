package qr

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ufguff/types"
	"github.com/ufguff/utils"
)

const (
	pathQr = "static/qr.png"
)

type Handler struct {
	store types.ICreateQr
}

func NewHandler(store types.ICreateQr) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/GetQR", h.handleQRRequest).Methods(http.MethodGet)
}

func (h *Handler) handleQRRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	urlV := vars["url"]
	sizeV := vars["size"]

	size, err := strconv.Atoi(sizeV)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	err = h.store.GetQRImage(urlV, size)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	http.ServeFile(w, r, pathQr)

	//delete
}
