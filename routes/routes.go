package routes

import (
	"net/http"

	controllers "github.com/jtonynet/golang-course-02/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
