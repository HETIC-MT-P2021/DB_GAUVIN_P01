package router

import (
	"log"
	"net/http"

	"github.com/jasongauvin/DB_GAUVIN_P01/controllers"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/employee", controllers.ShowEmployee)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
