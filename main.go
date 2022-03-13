package main

import(
	"log"
	"net/http"

	"todorestapi/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main() {
	r := mux.NewRouter()
	routes.RegisterTaskRoutes(r)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:8181", r))
}