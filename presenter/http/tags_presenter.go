package http

import (
	"encoding/json"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities"
	"net/http"
)

type tagsPresenter struct {
	Tags []entities.Tag `json:"tags"`
}

func TagsPresenter(w http.ResponseWriter, tags []entities.Tag) error {
	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(tagsPresenter{
		Tags: tags,
	})
	_, err := w.Write(b)
	return err
}

func TagPresenter(w http.ResponseWriter, tag entities.Tag) error {
	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(tag)
	_, err := w.Write(b)
	return err
}
