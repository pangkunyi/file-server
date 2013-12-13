package main

import (
	"config"
	"net/http"
	"regexp"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(config.MaxProcs)
	http.Handle("/public/silent_apk/", stripPrefix("/public/silent_apk/.*/apk/", FileServer(Dir(config.BaseDir))))
	if err := http.ListenAndServe(config.ServerAddr, nil); err != nil {
		panic(err)
	}
}

func stripPrefix(prefix string, h http.Handler) http.Handler {
	if prefix == "" {
		return h
	}
	re := regexp.MustCompile(prefix)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if p := re.ReplaceAllString(r.URL.Path, ""); len(p) < len(r.URL.Path) {
			r.URL.Path = p
			h.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	})
}
