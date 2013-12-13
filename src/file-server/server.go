package main

import (
	"config"
	"github.com/gorilla/mux"
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(config.MaxProcs)
	r := mux.NewRouter()
	http.Handle("/tmp/", http.StripPrefix("/tmp/", FileServer(Dir(config.BaseDir))))
	http.Handle("/", r)
	if err := http.ListenAndServe(config.ServerAddr, nil); err != nil {
		panic(err)
	}
}
