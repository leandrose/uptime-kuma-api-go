package http

import (
	"encoding/json"
	"net/http"
)

type errorPresenter struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ErrorPresenter(w http.ResponseWriter, headerCode int, err error) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(headerCode)
	b, e := json.Marshal(errorPresenter{
		Success: false,
		Message: err.Error(),
	})
	if e != nil {
		return e
	}
	_, e = w.Write(b)
	return e
}

func Error500Presenter(w http.ResponseWriter, err error) error {
	return ErrorPresenter(w, 500, err)
}

func Error400Presenter(w http.ResponseWriter, err error) error {
	return ErrorPresenter(w, 400, err)
}

func Error401Presenter(w http.ResponseWriter, err error) error {
	return ErrorPresenter(w, 401, err)
}

func Error402Presenter(w http.ResponseWriter, err error) error {
	return ErrorPresenter(w, 402, err)
}

func Error403Presenter(w http.ResponseWriter, err error) error {
	return ErrorPresenter(w, 403, err)
}

func Error404Presenter(w http.ResponseWriter, err error) error {
	return ErrorPresenter(w, 404, err)
}
