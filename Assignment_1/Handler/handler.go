package Handler

import (
	"fmt"
	"myapi/Config"
	"net/http"
	"sync"
)

type Repository struct {
	app *Config.App
}

var Repo *Repository

func NewRepo(app *Config.App) *Repository {
	return &Repository{
		app: app,
	}
}
func NewHandler(repo *Repository) {
	Repo = repo
}

func (m *Repository) GetViews(w http.ResponseWriter, r *http.Request) {
	req_id := r.URL.Query().Get(":id")
	var views int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		views = m.app.Get(req_id)
	}()
	wg.Wait()
	fmt.Fprintf(w, fmt.Sprintf("given video id:%v  Views:%v", req_id, views))

}

func (m *Repository) IncViews(w http.ResponseWriter, r *http.Request) {

	req_id := r.URL.Query().Get(":id")
	go m.app.Inc(req_id)
	fmt.Fprintf(w, fmt.Sprintf("Views Incremented for given video id:%v", req_id))
}
