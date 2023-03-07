package Config

import "sync"

type Video struct {
	Id    string `json:"id"`
	Views int    `json:"views"`
}

type App struct {
	V    []Video
	Size int
	Mp   map[string]int
	Mutex sync.Mutex
}

func (a *App) Get(id string) int {
	_, ok := a.Mp[id]
	a.Mutex.Lock()
	defer a.Mutex.Unlock()
	if !ok {
		vid := Video{
			Id:    id,
			Views: 0,
		}
		a.V = append(a.V, vid)
		a.Mp[id] = a.Size
		a.Size++
	}
	ind := a.Mp[id]
	return a.V[ind].Views
	

}

func (a *App) Inc(id string) {
	ind, ok := a.Mp[id]
	a.Mutex.Lock()
	defer a.Mutex.Unlock()
	if !ok {
		vid := Video{
			Id:    id,
			Views: 1,
		}
		a.V = append(a.V, vid)
		a.Mp[id] = a.Size
		a.Size++
	} else {
		a.V[ind].Views++
	}

}
