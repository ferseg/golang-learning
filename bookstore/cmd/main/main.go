package main

import (
	"github.com/ferseg/learning-go/bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
	"log"
	"net/http"
)

func main() {
    r := mux.NewRouter()
    // r.Headers("Content-Type", "application/json")
    routes.RegisterBookStoreRoutes(r)
    http.Handle("/", r)
    r.Use(contentTypeMiddleware)
    log.Fatal(http.ListenAndServe(":8080", r))
}

func contentTypeMiddleware(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        handler.ServeHTTP(w, r)
    })
}
