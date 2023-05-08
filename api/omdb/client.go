package omdb

import (
	"cinema/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type OMBD struct {
	client  http.Client
	apikey  string
	urlBase string
}

func NewOMBD(apikey string) *OMBD {
	client := http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://www.omdbapi.com/?apikey=%s", apikey)
	return &OMBD{client: client, apikey: apikey, urlBase: url}
}

func (o *OMBD) Read(imdbID string) models.FilmJson {
	url := o.urlBase + fmt.Sprintf("&i=%s", imdbID)
	r, err := o.client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	m := models.FilmJson{}
	err = json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		log.Fatal(err)
	}
	return m
}
