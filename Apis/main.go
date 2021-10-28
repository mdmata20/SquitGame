package main

import(
	"fmt"
	"net/http"
	"log"

	
	"github.com/gorilla/mux"
)


func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome API");
}





func main() {
	router := mux.NewRouter().StrictSlash(true)
	
	router.HandleFunc("/", indexRoute)

	log.Fatal(http.ListenAndServe(":3000", router))
	
}