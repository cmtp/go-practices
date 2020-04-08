package mux

import (
	"net/http"
)

type customeHandler func(http.ResponseWriter, *http.Request)

type MuxFacilito struct {
	rutas map[string]customeHandler //handlers
}

func (this *MuxFacilito) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if fn, ok := this.rutas[r.URL.Path]; ok {
		fn(w, r)
	} else {
		http.NotFound(w, r)
	}

}

func (this *MuxFacilito) AddFunc(ruta string, fn customeHandler) {
	this.rutas[ruta] = fn
}

func (this *MuxFacilito) AddHandle(ruta string, handle http.Handler) {
	this.rutas[ruta] = handle.ServeHTTP
}

func CreateMux() *MuxFacilito {
	newMap := make(map[string]customeHandler)
	return &MuxFacilito{newMap}
}
