package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) GenerateRoutes(middlewareFunctions ...func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)) (mainRouter *mux.Router) {

	mainRouter = mux.NewRouter()
	mainRouter.NotFoundHandler = http.HandlerFunc(h.NotFoundHandler)
	mainRouter.PathPrefix("/help").HandlerFunc(h.HelpHandler)
	mainRouter.PathPrefix("/scouts").HandlerFunc(h.ScoutsHandler).Methods("POST")
	mainRouter.PathPrefix("/pits").HandlerFunc(h.PitsHandler).Methods("POST")

	return

}
