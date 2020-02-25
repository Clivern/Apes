// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/clivern/apes/internal/app/controller"
	"github.com/clivern/apes/internal/app/module"

	"github.com/drone/envsubst"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	var configFile string
	var get string

	flag.StringVar(&configFile, "config", "config.prod.yml", "config")
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

	if get == "health" {
		return
	}

	configUnparsed, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(fmt.Sprintf(
			"Error while reading config file [%s]: %s",
			configFile,
			err.Error(),
		))
	}

	configParsed, err := envsubst.EvalEnv(string(configUnparsed))

	if err != nil {
		panic(fmt.Sprintf(
			"Error while parsing config file [%s]: %s",
			configFile,
			err.Error(),
		))
	}

	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer([]byte(configParsed)))

	if err != nil {
		panic(fmt.Sprintf(
			"Error while loading configs [%s]: %s",
			configFile,
			err.Error(),
		))
	}

	if viper.GetString("log.output") != "stdout" {
		fs := module.FileSystem{}
		dir, _ := filepath.Split(viper.GetString("log.output"))

		if !fs.DirExists(dir) {
			if _, err := fs.EnsureDir(dir, 777); err != nil {
				panic(fmt.Sprintf(
					"Directory [%s] creation failed with error: %s",
					dir,
					err.Error(),
				))
			}
		}

		if !fs.FileExists(viper.GetString("log.output")) {
			f, err := os.Create(viper.GetString("log.output"))
			if err != nil {
				panic(fmt.Sprintf(
					"Error while creating log file [%s]: %s",
					viper.GetString("log.output"),
					err.Error(),
				))
			}
			defer f.Close()
		}
	}

	r := mux.NewRouter()
	r.HandleFunc("/f/{path:.*}", controller.Proxy())
	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(viper.GetInt("app.port"))), r)
}
