package actions

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Serve() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomePage)
	http.Handle("/", r)
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatalln(err)
	}
}
