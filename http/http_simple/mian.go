package main

import (
	// "flag"
	"fmt"
	"io"
	"log"
	"net/http"

	// "github.com/golang/glog"
	"k8s.io/klog"
)

func main() {
	// flag.Set("v", "4")
	// glog.V(2).Info("Starting http server...")
	klog.Info("Starting http server...")
	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	user := r.URL.Query().Get("user")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranter]\n")
	}
	io.WriteString(w, "===================Details of the http request header:================\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
}

func healthz(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "ok")
}
