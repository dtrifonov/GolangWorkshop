package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/dtrifonov/GolangWorkshop/internal/diagnostics"
)

func main() {
	log.Print("Hello World")

	router := mux.NewRouter()
	router.HandleFunc("/", hello)

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	err := http.ListenAndServe(":8585", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}
