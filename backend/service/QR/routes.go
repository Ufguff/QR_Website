package qr

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ufguff/types"
	"github.com/ufguff/utils"
)

const (
	pathQr = "./static/qr.png"
)

type Handler struct {
	store types.ICreateQr
}

func NewHandler(store types.ICreateQr) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/GetQR", h.handleQRRequest).Methods(http.MethodPost)
}

func (h *Handler) handleQRRequest(w http.ResponseWriter, r *http.Request) {
	var payload *types.QR
	log.Println("begin")

	log.Print("parsing json")
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Println("validation")
	if err := utils.Validator.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Println("get qr")
	err := h.store.GetQRImage(*payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	http.ServeFile(w, r, pathQr)

	utils.WriteJSON(w, http.StatusOK, nil)

	os.Remove(pathQr)
}
