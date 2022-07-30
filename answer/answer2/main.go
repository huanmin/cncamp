package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"

	"k8s.io/klog"
)

func main() {
	klog.Info("Starting http server...")
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		klog.Error(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	klog.Info("用户请求IP: ", r.Host)

	version := os.Getenv("VERSION")
	w.Header().Add("VERSION", version)
	fmt.Printf("os version: %s \n", version)

	for k, v := range r.Header {
		klog.Info("用户请求日志:", k, "=", v)
		// w.Header().Set(k, fmt.Sprint(v))
		for _, v2 := range v {
			w.Header().Add(k, v2)
		}
	}

	// 04.记录日志并输出
	clientip := getCurrentIP(r)
	//fmt.Println(r.RemoteAddr)
	klog.Info("Success! Response code: %d", 200)
	klog.Info("Success! clientip: %s", clientip)
}

func healthz(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "ok")
	w.WriteHeader(200)
}

func getCurrentIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
