package handlers

import (
	"net/http"
	"time"
)

func HelloWorldHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte("Hello World"))
	time.Sleep(5 * time.Second)
}
