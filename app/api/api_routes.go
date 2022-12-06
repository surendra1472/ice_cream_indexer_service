package api

import (
	"github.com/gorilla/mux"
	"ic-indexer-service/app/api/controller"
	"ic-indexer-service/ice_cream_middleware"
	"log"
	"net/http"
)

const (
	getIcecream    = "GetIcecream"
	putIcecream    = "PutIcecream"
	deleteIcecream = "deleteIcecream"
)

func GetRoutes() *mux.Router {
	r := mux.NewRouter()

	v1UnAuthenticatedRouter := r.PathPrefix("/internal/v1").Subrouter()
	v1UnAuthenticatedRouter.HandleFunc("/health_check", controller.GetHeartBeat).Methods(http.MethodGet).Name("GetHeartBeat")

	v1Router := r.PathPrefix("/api/v1").Subrouter()
	v1RouterInternal := r.PathPrefix("/api/v1").Subrouter()

	//GET -- READ
	v1Router.HandleFunc("/icecream", controller.GetIcecream).Methods(http.MethodGet).Name(getIcecream)

	//PUT
	v1RouterInternal.HandleFunc("/icecream", controller.UpdateIcecream).Methods(http.MethodPut).Name(putIcecream)

	//DELETE
	v1RouterInternal.HandleFunc("/icecream", controller.DeleteIcecream).Methods(http.MethodDelete).Name(deleteIcecream)

	addMiddlewares(v1Router, true)
	addMiddlewares(v1RouterInternal, false)

	return r
}

func addMiddlewares(routes *mux.Router, isAuthenticated bool) {
	log.Print("Adding Middlewares")
	routes.Use(ice_cream_middleware.GenerateRequestIdHandler)

	if isAuthenticated {
		routes.Use(ice_cream_middleware.TokenHandler)
		routes.Use(ice_cream_middleware.AuthMiddleware)
	}

}
