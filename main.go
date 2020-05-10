package main

import (
	"net/http"

	"github.com/bwangelme/golang-elastic-api/app/Routers"
	"github.com/bwangelme/golang-elastic-api/app/Utils"
	"github.com/gorilla/mux"
	//"fmt"
)

func main() {
	/*
		Routing using mux
	*/

	Utils.GetNumCpu()
	r := mux.NewRouter()
	r.HandleFunc("/set", Routers.SetHandler)
	r.HandleFunc("/get", Routers.GetHandler)
	r.HandleFunc("/map", Routers.MappingHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8000", r)

}
