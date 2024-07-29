package main

import (
	"log"
	"net/http"
	"recruitment-system/internal/handlers"
	"github.com/gorilla/mux"
	//"recruitment-system/internal/db"
	//"database/sql"
    //"github.com/lib/pq"
)

func main() {
	//db.Init()
    r := mux.NewRouter()
    handlers.RegisterRoutes(r)
    log.Println("Server started at :8080")
    http.ListenAndServe(":8080", r)
}
