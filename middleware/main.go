package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := "localhost:4000"

	mux := http.NewServeMux()
	muxLogged := http.NewServeMux()
	muxLogged.HandleFunc("/v1/hello1", HelloHandler1)
	muxLogged.HandleFunc("/v1/hello2", HelloHandler2)

	// /v1/hello3 will not log because it is on mux
	// but /v1/ will b/c of the Adapter
	mux.HandleFunc("/v1/hello3", HelloHandler3)
	logger := log.New(os.Stdout, "", log.LstdFlags)
	mux.Handle("/v1/", Adapt(muxLogged, logRequests(logger), throttleRequests(2, time.Minute)))

	fmt.Printf("listening at %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
