// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/clivern/apes/internal/app/module"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Proxy controller
func Proxy(c *gin.Context) {
	path := c.Param("path")
	target := "https://httpbin.org/status/200"

	logger, _ := module.NewLogger(
		viper.GetString("log.level"),
		viper.GetString("log.format"),
		[]string{viper.GetString("log.output")},
	)

	defer func() {
		_ = logger.Sync()
	}()

	url, err := url.Parse(target)

	if err != nil {
		logger.Error(fmt.Sprintf(
			`Error: %s`,
			err.Error(),
		), zap.String("CorrelationId", c.Request.Header.Get("X-Correlation-ID")))

		c.Status(http.StatusInternalServerError)
		return
	}

	logger.Info(fmt.Sprintf(
		`Proxy %s -> %s`,
		path,
		target,
	), zap.String("CorrelationId", c.Request.Header.Get("X-Correlation-ID")))

	proxy := httputil.NewSingleHostReverseProxy(url)

	//c.Request.Host = url.Host
	proxy.ServeHTTP(c.Writer, c.Request)
}
