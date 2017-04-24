package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func logReq(r *http.Request) {
	log.Println(r.Method, r.URL.Path)
}

//closure to handle logs for all handlers
//have to use in mux (main.go) and wrap each handler
func logReqs(hfn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		//logs the time it takes to process the request
		start := time.Now()
		hfn(w, r)
		fmt.Printf("%v\n", time.Since(start))
	}
}

//wrap entire mux approach
func logRequests(logger *log.Logger) Adapter {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Printf("%s %s", r.Method, r.URL.Path)
			start := time.Now()
			handler.ServeHTTP(w, r)
			logger.Printf("%v\n", time.Since(start))
		})
	}
}
