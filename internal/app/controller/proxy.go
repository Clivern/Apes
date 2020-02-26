// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

// Proxy controller
func Proxy(upstream string, _ string, _ string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		remote, err := url.Parse(upstream)

		if err != nil {
			log.Fatalf("Unexpected Error: %s", err.Error())
		}

		log.Printf(
			"Proxy %s to %s%s",
			r.URL.Path,
			upstream,
			r.URL.Path,
		)

		proxy := httputil.NewSingleHostReverseProxy(remote)

		r.URL.Path = mux.Vars(r)["path"]
		proxy.ServeHTTP(w, r)
	}
}
