package main

import (
	"crypto/tls"
	"net/http"

	"github.com/peterzam/OutlineBot/controller"
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func main() {
	controller.StartBot()
}
