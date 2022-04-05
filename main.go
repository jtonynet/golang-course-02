package main

import (
	"net/http"

	routes "github.com/jtonynet/golang-course-02/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8001", nil)
}
