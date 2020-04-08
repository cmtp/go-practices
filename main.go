package main

import (
	"fmt"
	"log"
	"net/http"
)

type customeHandler func(http.ResponseWriter, *http.Request)

type MuxFacilito struct {
	rutas map[string]customeHandler //handlers
}

func (this *MuxFacilito) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn := this.rutas[r.URL.Path]
	fn(w, r)
}

func (this *MuxFacilito) AddMux(ruta string, fn customeHandler) {
	this.rutas[ruta] = fn
}

func main() {
	newMap := make(map[string]customeHandler)
	mux := &MuxFacilito{rutas: newMap}
	mux.AddMux("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola desde una funcion anonima")
	})

	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
