package qr

import (
	"log"
	"net/http"

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

	if err := utils.ParseJSON(r, &payload); err != nil {
		log.Println(err, "parse")
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validator.Struct(payload); err != nil {
		log.Println(err, "validate")
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Println(payload.Url)
	log.Println(payload.Size)

	err := h.store.GetQRImage(payload.Url, payload.Size)
	if err != nil {
		log.Println(err, "getQR")
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	http.ServeFile(w, r, pathQr)

	//delete
	utils.WriteJSON(w, http.StatusOK, "qr send")
}
