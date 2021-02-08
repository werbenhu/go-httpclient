// Copyright 2014-2019 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package main

import (
	"fmt"
	"github.com/ddliu/go-httpclient"
	"net/http"
)

const (
	USERAGENT       = "my awsome httpclient"
	TIMEOUT         = 30
	CONNECT_TIMEOUT = 5
	SERVER          = "http://127.0.0.1/endpoint"
)

func main() {
	httpclient.Defaults(httpclient.Map{
		"opt_useragent":   USERAGENT,
		"opt_timeout":     TIMEOUT,
		"Accept-Encoding": "gzip,deflate,sdch",
	})

	res, _ := httpclient.
		WithHeader("Accept-Language", "en-us").
		WithCookie(&http.Cookie{
			Name:  "name",
			Value: "github",
		}).
		WithHeader("Referer", "http://google.com").
		PostJson(SERVER, "abc")


	fmt.Println("Cookies:")
	for k, v := range httpclient.CookieValues(SERVER) {
		fmt.Println(k, ":", v)
	}

	type Obj struct {
		Header map[string]string `json:"header"`
		Payload map[string]string `json:"payload"`
	}

	fmt.Println("Response:")
	var obj Obj

	err := res.Unmarshal(&obj)

	fmt.Println(err)
	fmt.Println(obj)
}
