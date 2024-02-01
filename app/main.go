package main

import (
	"encoding/json"
	"net"
	"net/http"
	"time"
)

type Response struct {
	Timestamp string `json:"timestamp"`
	IP        string `json:"ip"`
}

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func handler(w http.ResponseWriter, r *http.Request) {
	ip, _, _ := net.SplitHostPort(getIP(r))
	resp := Response{
		Timestamp: time.Now().Format(time.RFC3339),
		IP:        ip,
	}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
