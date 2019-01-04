package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	env := os.Getenv("ENVIRONMENT")
	port := 8081
	flag.Parse()
	log.Printf("Backend environment: %s\n", env)

	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received environment check from %s", r.RemoteAddr)
		fmt.Fprintf(w, "%s", env)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request from %s at %s", r.RemoteAddr, r.URL.EscapedPath())
		i := InstanceMetadata{}
		i.Populate(env)
		raw, _ := httputil.DumpRequest(r, true)
		i.LBRequest = string(raw)
		resp, _ := json.Marshal(i)
		fmt.Fprintf(w, "%s", resp)
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received health check from %s", r.RemoteAddr)
		w.WriteHeader(http.StatusOK)
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}
