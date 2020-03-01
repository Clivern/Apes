// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// Proxy controller
func Proxy(upstream string, failRate string, latency string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(time.Now().UnixNano())
		remote, err := url.Parse(upstream)

		if err != nil {
			log.Fatalf("Unexpected Error: %s", err.Error())
		}

		log.Printf(
			"Incoming request to Proxy %s to %s%s",
			r.URL.Path,
			upstream,
			r.URL.Path,
		)

		failCount, _ := strconv.Atoi(strings.Replace(failRate, "%", "", -1))

		if rand.Intn(100) < failCount {
			log.Printf(
				"Fail to Proxy %s to %s%s",
				r.URL.Path,
				upstream,
				r.URL.Path,
			)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		latencySeconds, _ := strconv.Atoi(strings.Replace(latency, "s", "", -1))

		log.Printf("Hold on for %s", latency)

		time.Sleep(time.Duration(latencySeconds) * time.Second)

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
