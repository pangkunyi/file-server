package main

import (
	"config"
	"log"
	"net/http"
	_ "net/http/pprof"
	"path/filepath"
	"regexp"
	"runtime"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe(config.C.DebugAddr, nil))
	}()
	runtime.GOMAXPROCS(config.C.MaxProcs)
	for _, rule := range config.C.Rules {
		if rule.Cached {
			http.Handle(rule.Pattern, stripPrefix(rule.Strip, FileServer(Dir(rule.Dir))))
		} else {
			http.Handle(rule.Pattern, stripPrefix(rule.Strip, http.FileServer(http.Dir(rule.Dir))))
		}
	}

	srv := &http.Server{Addr: config.C.ServerAddr, ReadTimeout: config.C.ReadTimeout, WriteTimeout: config.C.WriteTimeout}
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
			w.Header().Set("filename", filepath.Base(p))
			h.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	})
}
