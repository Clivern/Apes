// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

func Proxy() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		target := "https://httpbin.org"
		remote, err := url.Parse(target)

		if err != nil {
			panic(err)
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)

		r.URL.Path = mux.Vars(r)["path"]
		proxy.ServeHTTP(w, r)
	}
}
