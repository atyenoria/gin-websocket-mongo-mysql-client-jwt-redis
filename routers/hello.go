package routers

import (
	"sample/examples/etc-api/golang-jwt-authentication-api-sample/controllers"
	"sample/examples/etc-api/golang-jwt-authentication-api-sample/core/authentication"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle("/test/hello",
		negroni.New(

				negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")

	return router
}
