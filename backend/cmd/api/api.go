package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	qr "github.com/ufguff/service/QR"
)

type ApiServer struct {
	addr string
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{addr: addr}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()

	subrouter := router.PathPrefix("api/v1").Subrouter()

	qrStore := qr.Store{}
	qrHandler := qr.NewHandler(&qrStore)
	qrHandler.RegisterRoutes(subrouter)

	c := cors.Default()

	handler := c.Handler(subrouter)

	log.Println("Listeting on " + s.addr)
	return http.ListenAndServe(s.addr, handler)
}
