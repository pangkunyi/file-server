package main

import (
	"github.com/pangkunyi/plum/logs"
	"net/http"
	_ "net/http/pprof"
	"regexp"
	"runtime"
)

var (
	ALog *logs.Logger
	MLog *logs.Logger
)

func init() {
	MLog = logs.NewLogger(C.MainLogFile, false)
	ALog = logs.NewLogger(C.AccessLogFile, true)
	MLog.Printf("config:%#v\n", C)
}

func main() {
	runtime.GOMAXPROCS(C.MaxProcs)
	go func() {
		MLog.Printf("server failed:%v\n", http.ListenAndServe(C.DebugAddr, nil))
	}()
	for _, rule := range C.Rules {
		if rule.Cached {
			http.Handle(rule.Pattern, stripPrefix(rule.Strip, FileServer(Dir(rule.Dir))))
		} else {
			http.Handle(rule.Pattern, stripPrefix(rule.Strip, http.FileServer(http.Dir(rule.Dir))))
		}
	}

	if len(C.ServerAddrs) < 1 {
		MLog.Fatal("server addr not set!\n")
	}
	for i := 0; i < len(C.ServerAddrs)-1; i++ {
		go func(idx int) {
			startServer(C.ServerAddrs[idx])
		}(i)
	}
	startServer(C.ServerAddrs[len(C.ServerAddrs)-1])
}

func startServer(serverAddr string) {
	srv := &http.Server{Addr: serverAddr, ReadTimeout: C.ReadTimeout, WriteTimeout: C.WriteTimeout}
	MLog.Printf("server failed:%v\n", srv.ListenAndServe())
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
