package main

import (
	"net/http"

	"github.com/tranqhuy46/go-webservices/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3001", nil)
}
