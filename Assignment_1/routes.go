package main

import (
	"myapi/Handler"
	"net/http"

	"github.com/bmizerany/pat"
)

func route() http.Handler {
	mux := pat.New()
	mux.Get("/:id", http.HandlerFunc(Handler.Repo.GetViews))
	mux.Post("/:id", http.HandlerFunc(Handler.Repo.IncViews))
	return mux
}
