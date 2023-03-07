package main

import (
	"log"
	"myapi/Config"
	"myapi/Handler"
	"net/http"
)

func main() {

	app := Config.App{
		Mp:   make(map[string]int),
		V:    make([]Config.Video, 0),
		Size: 0,
	}
	repo := Handler.NewRepo(&app)
	Handler.NewHandler(repo)
	mux := route()
	http.Handle("/", mux)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		log.Println("Port running on :12345")
	}

}
