package main

import (
	"fmt"
	"log"

	"server_mux/handler"
	"server_mux/sqlite"
)
import "net/http"

var ROUTING_TABLE = map[string]http.Handler{
	"/a":    http.HandlerFunc(handler.HandleA),
	"/b":    http.HandlerFunc(handler.HandleB),
	"/test": http.HandlerFunc(handler.HandleTest),
	"/":     http.HandlerFunc(handleIndex),
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index Page")
}

func initRoutingTable(mux *http.ServeMux) {
	for pattern, handlerFunc := range ROUTING_TABLE {
		mux.Handle(pattern, handlerFunc)
	}
}

func initDatabase() {
	sqlite.Init()
}

func initMuxServer() *http.ServeMux {
	mux := http.NewServeMux()
	initRoutingTable(mux)
	initDatabase()
	return mux
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", initMuxServer()))
}
