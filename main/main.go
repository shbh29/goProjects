package main

import (
	"localhost/shbh/webservice/controller"
	"net/http"
)

func main() {
	controller.RegisterController()
	http.ListenAndServe(":3000", nil)
}
