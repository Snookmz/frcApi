package handler

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

func (h *Handler) GenerateRoutes(middlewareFunctions ...func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)) (mainRouter *mux.Router) {

	mainRouter = mux.NewRouter()
	mainRouter.NotFoundHandler = http.HandlerFunc(h.NotFoundHandler)
	mainRouter.PathPrefix("/help").HandlerFunc(h.HelpHandler)
	mainRouter.PathPrefix("/login").HandlerFunc(h.LoginHandler).Methods("POST")

	var authRoutes *mux.Router
	authRoutes = mux.NewRouter()
	authRoutes.NotFoundHandler = http.HandlerFunc(h.NotFoundHandler)


	// setting up middle ware for sub router
	n := negroni.New()
	for i := 0; i < len(middlewareFunctions); i++ {
		n.Use(negroni.HandlerFunc(middlewareFunctions[i]))
	}
	n.Use(negroni.Wrap(authRoutes))
	mainRouter.PathPrefix("/policy").Handler(n)
	mainRouter.PathPrefix("/user").Handler(n)
	mainRouter.PathPrefix("/data").Handler(n)

	policySub := authRoutes.PathPrefix("/policy").Subrouter()
	userSub := authRoutes.PathPrefix("/user").Subrouter()
	dataSub := authRoutes.PathPrefix("/data").Subrouter()

	policySub.Path("/list").HandlerFunc(h.PolicyListHandler).Methods("POST")
	policySub.Path("/get/active").HandlerFunc(h.PolicyGetActiveHandler).Methods("POST")
	policySub.Path("/get/id").HandlerFunc(h.PolicyGetByPolicyIdHandler).Methods("POST")
	userSub.Path("/logout").HandlerFunc(h.LogoutHandler).Methods("POST")
	userSub.Path("/check/session").HandlerFunc(h.CheckSessionHandler).Methods("POST")

	dataSub.Path("").HandlerFunc(h.GetDataHandler).Methods("POST")
	dataSub.Path("/activity").HandlerFunc(h.GetActivityHandler).Methods("POST")
	dataSub.Path("/top-n").HandlerFunc(h.GetTopNHandler).Methods("POST")
	dataSub.Path("/top-classes").HandlerFunc(h.GetTopClassesHandler).Methods("POST")


	return

}
