package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/ruprict/loccasions-go/api"
	"github.com/ruprict/loccasions-go/app"
	_ "github.com/ruprict/loccasions-go/db"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	db, err := gorm.Open("postgres", "user=postgres dbname=loccasions_development sslmode=disable")
	if err != nil {
		panic(err)
	}
	context := app.Context{&db}

	//http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	//r.PathPrefix("/public/").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	r.HandleFunc("/api/users", context.Handler(api.UsersCreateHandler)).Methods("POST")
	r.HandleFunc("/api/users", context.Handler(api.UsersIndexHandler)).Methods("GET")
	r.HandleFunc("/api/users/{userId}", context.Handler(api.UsersUpdateHandler)).Methods("PUT", "PATCH")
	r.HandleFunc("/api/users/{userId}/loccasions", context.Handler(api.LoccasionsIndexHandler)).Methods("GET")
	r.HandleFunc("/api/users/{userId}/loccasions", context.Handler(api.LoccasionsCreateHandler)).Methods("POST")
	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":8080")
}

func HomeHandler(rw http.ResponseWriter, req *http.Request) {

}
