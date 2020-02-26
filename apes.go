// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/clivern/apes/internal/app/controller"

	"github.com/gorilla/mux"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	var get string
	var port int
	var upstream string
	var failRate string
	var latency string

	flag.IntVar(&port, "port", 8080, "port")
	flag.StringVar(&upstream, "upstream", "https://httpbin.org", "upstream")
	flag.StringVar(&failRate, "failRate", "10%", "failRate")
	flag.StringVar(&latency, "latency", "0s", "latency")
	flag.StringVar(&get, "get", "", "get")
	flag.Parse()

	if get == "release" {
		fmt.Println(
			fmt.Sprintf(
				`Apes Version %v Commit %v, Built @%v`,
				version,
				commit,
				date,
			),
		)
		return
	}

	log.Printf(
		"Starting apes chaos reverse proxy port=%d upstream=%s failRate=%s latency=%s",
		port,
		upstream,
		failRate,
		latency,
	)

	r := mux.NewRouter()
	r.HandleFunc("/{path:.*}", controller.Proxy(upstream, failRate, latency))
	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
