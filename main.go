package main

import (
	"flag"
	"net/http"
	"time"

	api "apis/controller"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", hello)
	router.HandleFunc("/as", api.Register)
	router.HandleFunc("/login", api.Login)
	router.HandleFunc("/getusers", api.IsAuthorized(api.GetUsers))
	router.HandleFunc("/update", api.UpdateUser)
	router.HandleFunc("/delete/{id}", api.DeleteUser)

	startServer := flag.Bool("startserver", false, "start the server on port 8080")
	flag.Parse()
	if *startServer {
		s := &http.Server{
			Addr:         ":8080",
			Handler:      router,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		log.Fatal(s.ListenAndServe())
	}

}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
