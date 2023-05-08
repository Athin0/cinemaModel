package models

import (
	"log"
	"strconv"
	"strings"
	"time"
)

type FilmJson struct {
	FilmID      string `json:"imdbID"`   // film_id
	Title       string `json:"Title"`    // title
	Description string `json:"Plot"`     // description
	Duration    string `json:"Runtime"`  // duration
	ReleaseDate string `json:"Year"`     // release_date
	AgeRating   string `json:"Rated"`    // age_rating
	Country     string `json:"Country"`  // country_id
	Producer    string `json:"Director"` // producer_id
	Genre       string `json:"Genre"`    // genre_id
	Actors      string `json:"Actors"`   // actor_id
}

func (f *FilmJson) GetCountries() []string {
	arr := make([]string, 0)
	for _, s := range strings.Split(f.Country, ",") {
		arr = append(arr, strings.TrimSpace(s))
	}
	return arr
}

func (f *FilmJson) GetActors() []string {
	arr := make([]string, 0)
	for _, s := range strings.Split(f.Actors, ",") {
		arr = append(arr, strings.TrimSpace(s))
	}
	return arr
}
func (f *FilmJson) GetGenres() []string {
	arr := make([]string, 0)
	for _, s := range strings.Split(f.Genre, ",") {
		arr = append(arr, strings.TrimSpace(s))
	}
	return arr
}

func (f *FilmJson) GetDate() time.Time {
	//date := strings.Split(f.ReleaseDate, " ")
	year, err := strconv.Atoi(f.ReleaseDate)
	if err != nil {
		log.Fatal(err)
	}
	return time.Date(year, 0, 0, 0, 0, 0, 0, time.UTC)
}

func (f *FilmJson) GetDuration() int {
	minutes, err := strconv.Atoi(strings.Split(f.Duration, " ")[0])
	if err != nil {
		log.Fatal(err)
	}
	return minutes
}
func Month(s string) time.Month {
	switch s {
	case "Dec":
		return time.December
	case "Jan":
		return time.January
	case "May":
		return time.May
	case "Jun":
		return time.June
	case "Aug":
		return time.August
	case "Oct":
		return time.October
	}
	return time.Month(0)
}
