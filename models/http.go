package models

import (
	"github.com/gorilla/mux"
	"net/http"
)

var (
	Multiplexer *http.ServeMux
	Router *mux.Router
)
