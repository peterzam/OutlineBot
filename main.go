package main

import (
	"crypto/tls"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/peterzam/OutlineBot/controller"
)

func init() {
	godotenv.Load("env.list")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func main() {
	controller.StartBot()
}
