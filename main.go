package main

import (
	"fmt"
	"log"
	"net/http"

	dummydb "com.web.level5/dummyDB"
	"github.com/gorilla/mux"
)

var port = ":8000"

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
	fmt.Println("Method used:", r.Method)
}

func handlerRequest() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink)
	router.HandleFunc("/dummys", dummydb.CreateEvent).Methods("POST")
	router.HandleFunc("/dummy", dummydb.GetAllEvents).Methods("GET")
	router.HandleFunc("/dummy/{id}", dummydb.GetOneEvent).Methods("GET")
	router.HandleFunc("/dummy/{id}", dummydb.UpdateEvent).Methods("PATCH")
	router.HandleFunc("/dummy/{id}", dummydb.DeleteEvent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(port, router))
}

func main() {
	handlerRequest()
}
