// Copyright 2014 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	_ "expvar"
	"net/http"

	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
)

// NewRouter is a function that returns a new configured Router
// It add the necessary middlewares
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	if Conf.Expvar {
		if !Conf.WasProvidedUserAndPassword() {
			panic("Your are exporting debug variables and looks like you forget to define an User and a Password")
		}

		router.Handle("/debug/vars", httpauth.SimpleBasicAuth(Conf.User, Conf.Password)(http.DefaultServeMux))
	}

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		if route.RequiresRestAuth {
			handler = RestAuthenticationHandler(handler)
			handler = RestCheckAppDisabledHandler(handler)
		}

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	return router
}
