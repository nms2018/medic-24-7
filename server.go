package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Server will start at http://127.0.0.1:8000 or 8001/")

	connectDatabse()

	route := mux.NewRouter()

	addApproutes(route)

	log.Fatal(http.ListenAndServe(":8001", route))
}
