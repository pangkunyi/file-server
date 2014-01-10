package main

import (
	"config"
	"log"
	"net/http"
	_ "net/http/pprof"
	"regexp"
	"runtime"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe(config.DebugAddr, nil))
	}()
	runtime.GOMAXPROCS(config.MaxProcs)
	http.Handle("/public/silent_apk/", stripPrefix("/public/silent_apk/.*/apk/", FileServer(Dir(config.BaseDir))))
	srv := &http.Server{Addr: config.ServerAddr, ReadTimeout: config.ReadTimeout, WriteTimeout: config.WriteTimeout}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
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
