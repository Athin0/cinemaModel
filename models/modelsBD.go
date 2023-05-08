package models

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"time"
)

// Consumer represents public.consumer
type Consumer struct {
	ConsumerID int    // consumer_id
	Login      string // login
	Password   string // password
}

// Create inserts the Consumer to the database.
func (r *Consumer) Create(db Queryer) error {
	return r.CreateContext(context.Background(), db)
}

// GetConsumerByPk select the Consumer from the database.
func GetConsumerByPk(db Queryer, pk0 int) (*Consumer, error) {
	return GetConsumerByPkContext(context.Background(), db, pk0)
}

// CreateContext inserts the Consumer to the database.
func (r *Consumer) CreateContext(ctx context.Context, db Queryer) error {
	err := db.QueryRowContext(ctx,
		`INSERT INTO consumer (login, password) VALUES ($1, $2) RETURNING consumer_id`,
		&r.Login, &r.Password).Scan(&r.ConsumerID)
	if err != nil {
		return errors.Wrap(err, "failed to insert consumer")
	}
	return nil
}

// GetConsumerByPkContext select the Consumer from the database.
func GetConsumerByPkContext(ctx context.Context, db Queryer, pk0 int) (*Consumer, error) {
	var r Consumer
	err := db.QueryRowContext(ctx,
		`SELECT consumer_id, login, password FROM consumer WHERE consumer_id = $1`,
		pk0).Scan(&r.ConsumerID, &r.Login, &r.Password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select consumer")
	}
	return &r, nil
}

// Country represents public.country
type Country struct {
	CountryID int    // country_id
	Name      string // name
}

// Create inserts the Country to the database.
func (r *Country) Create(db Queryer) error {
	return r.CreateContext(context.Background(), db)
}

// GetCountryByPk select the Country from the database.
func GetCountryByPk(db Queryer, pk0 int) (*Country, error) {
	return GetCountryByPkContext(context.Background(), db, pk0)
}

// CreateContext inserts the Country to the database.
func (r *Country) CreateContext(ctx context.Context, db Queryer) error {
	err := db.QueryRowContext(ctx,
		`INSERT INTO country (name) VALUES ($1) RETURNING country_id`,
		&r.Name).Scan(&r.CountryID)
	if err != nil {
		return errors.Wrap(err, "failed to insert country")
	}
	return nil
}

// GetCountryByPkContext select the Country from the database.
func GetCountryByPkContext(ctx context.Context, db Queryer, pk0 int) (*Country, error) {
	var r Country
	err := db.QueryRowContext(ctx,
		`SELECT country_id, name FROM country WHERE country_id = $1`,
		pk0).Scan(&r.CountryID, &r.Name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select country")
	}
	return &r, nil
}

// Film represents public.film
type Film struct {
	FilmID      int       // film_id
	Title       string    // title
	Description string    // description
	Duration    int       // duration
	ReleaseDate time.Time // release_date
	AgeRating   string    // age_rating
	ProducerID  int       // producer_id
}

// Create inserts the Film to the database.
func (r *Film) Create(db Queryer) error {
	return r.CreateContext(context.Background(), db)
}

// GetFilmByPk select the Film from the database.
func GetFilmByPk(db Queryer, pk0 int) (*Film, error) {
	return GetFilmByPkContext(context.Background(), db, pk0)
}

// CreateContext inserts the Film to the database.
func (r *Film) CreateContext(ctx context.Context, db Queryer) error {
	err := db.QueryRowContext(ctx,
		`INSERT INTO film (title, description, duration, release_date, age_rating, producer_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING film_id`,
		&r.Title, &r.Description, &r.Duration, &r.ReleaseDate, &r.AgeRating, &r.ProducerID).Scan(&r.FilmID)
	if err != nil {
		return errors.Wrap(err, "failed to insert film")
	}
	return nil
}

// GetFilmByPkContext select the Film from the database.
func GetFilmByPkContext(ctx context.Context, db Queryer, pk0 int) (*Film, error) {
	var r Film
	err := db.QueryRowContext(ctx,
		`SELECT film_id, title, description, duration, release_date, age_rating, producer_id FROM film WHERE film_id = $1`,
		pk0).Scan(&r.FilmID, &r.Title, &r.Description, &r.Duration, &r.ReleaseDate, &r.AgeRating, &r.ProducerID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select film")
	}
	return &r, nil
}

// FilmActor represents public.film_actor
type FilmActor struct {
	PersonID int // person_id
	FilmID   int // film_id
}

// Create inserts the FilmActor to the database.
func (r *FilmActor) Create(db Queryer) error {
	return r.CreateContext(context.Background(), db)
}

// GetFilmActorByPk select the FilmActor from the database.
func GetFilmActorByPk(db Queryer, pk0 int, pk1 int) (*FilmActor, error) {
	return GetFilmActorByPkContext(context.Background(), db, pk0, pk1)
}

// CreateContext inserts the FilmActor to the database.
func (r *FilmActor) CreateContext(ctx context.Context, db Queryer) error {
	_, err := db.ExecContext(ctx,
		`INSERT INTO film_actor (person_id, film_id) VALUES ($1, $2)`,
		&r.PersonID, &r.FilmID)
	if err != nil {
		return errors.Wrap(err, "failed to insert film_actor")
	}
	return nil
}

// GetFilmActorByPkContext select the FilmActor from the database.
func GetFilmActorByPkContext(ctx context.Context, db Queryer, pk0 int, pk1 int) (*FilmActor, error) {
	var r FilmActor
	err := db.QueryRowContext(ctx,
		`SELECT person_id, film_id FROM film_actor WHERE person_id = $1 AND film_id = $2`,
		pk0, pk1).Scan(&r.PersonID, &r.FilmID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select film_actor")
	}
	return &r, nil
}

// FilmCountry represents public.film_country
type FilmCountry struct {
	CountryID int // country_id
	FilmID    int // film_id
}

// Create inserts the FilmCountry to the database.
func (r *FilmCountry) Create(db Queryer) error {
	return r.CreateContext(context.Background(), db)
}

// GetFilmCountryByPk select the FilmCountry from the database.
func GetFilmCountryByPk(db Queryer, pk0 int, pk1 int) (*FilmCountry, error) {
	return GetFilmCountryByPkContext(context.Background(), db, pk0, pk1)
}

// CreateContext inserts the FilmCountry to the database.
func (r *FilmCountry) CreateContext(ctx context.Context, db Queryer) error {
	_, err := db.ExecContext(ctx,
		`INSERT INTO film_country (country_id, film_id) VALUES ($1, $2)`,
		&r.CountryID, &r.FilmID)
	if err != nil {
		return errors.Wrap(err, "failed to insert film_country")
	}
	return nil
}

// GetFilmCountryByPkContext select the FilmCountry from the database.
func GetFilmCountryByPkContext(ctx context.Context, db Queryer, pk0 int, pk1 int) (*FilmCountry, error) {
	var r FilmCountry
	err := db.QueryRowContext(ctx,
		`SELECT country_id, film_id FROM film_country WHERE country_id = $1 AND film_id = $2`,
		pk0, pk1).Scan(&r.CountryID, &r.FilmID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select film_country")
	}
	return &r, nil
}

// FilmGenre represents public.film_genre
type FilmGenre struct {
	GenreID int // genre_id
	FilmID  int // film_id
}

// Create inserts the FilmGenre to the database.
func (r *FilmGenre) Create(db Queryer) error {
	return r.CreateContext(context.Background(), db)
}

// GetFilmGenreByPk select the FilmGenre from the database.
func GetFilmGenreByPk(db Queryer, pk0 int, pk1 int) (*FilmGenre, error) {
	return GetFilmGenreByPkContext(context.Background(), db, pk0, pk1)
}

// CreateContext inserts the FilmGenre to the database.
func (r *FilmGenre) CreateContext(ctx context.Context, db Queryer) error {
	_, err := db.ExecContext(ctx,
		`INSERT INTO film_genre (genre_id, film_id) VALUES ($1, $2)`,
		&r.GenreID, &r.FilmID)
	if err != nil {
		return errors.Wrap(err, "failed to insert film_genre")
	}
	return nil
}

// GetFilmGenreByPkContext select the FilmGenre from the database.
func GetFilmGenreByPkContext(ctx context.Context, db Queryer, pk0 int, pk1 int) (*FilmGenre, error) {
	var r FilmGenre
	err := db.QueryRowContext(ctx,
		`SELECT genre_id, film_id FROM film_genre WHERE genre_id = $1 AND film_id = $2`,
		pk0, pk1).Scan(&r.GenreID, &r.FilmID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select film_genre")
	}
	return &r, nil
}

// Format represents public.format
type Format struct {
	FormatID int    // format_id
	Name     string // name
}

// Create inserts the Format to the database.
func (r *Format) Create(db Queryer) error {
	return r.CreateContext(context.Background(), db)
}

// GetFormatByPk select the Format from the database.
func GetFormatByPk(db Queryer, pk0 int) (*Format, error) {
	return GetFormatByPkContext(context.Background(), db, pk0)
}

// CreateContext inserts the Format to the database.
func (r *Format) CreateContext(ctx context.Context, db Queryer) error {
	err := db.QueryRowContext(ctx,
		`INSERT INTO format (name) VALUES ($1) RETURNING format_id`,
		&r.Name).Scan(&r.FormatID)
	if err != nil {
		return errors.Wrap(err, "failed to insert format")
	}
	return nil
}

// GetFormatByPkContext select the Format from the database.
func GetFormatByPkContext(ctx context.Context, db Queryer, pk0 int) (*Format, error) {
	var r Format
	err := db.QueryRowContext(ctx,
		`SELECT format_id, name FROM format WHERE format_id = $1`,
		pk0).Scan(&r.FormatID, &r.Name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select format")
	}
	return &r, nil
}

// Genre represents public.genre
type Genre struct {
	GenreID int    // genre_id
	Name    string // name
}

// Create inserts the Genre to the database.
func (r *Genre) Create(db Queryer) error {
	return r.CreateContext(context.Background(), db)
}

// GetGenreByPk select the Genre from the database.
func GetGenreByPk(db Queryer, pk0 int) (*Genre, error) {
	return GetGenreByPkContext(context.Background(), db, pk0)
}

// CreateContext inserts the Genre to the database.
func (r *Genre) CreateContext(ctx context.Context, db Queryer) error {
	err := db.QueryRowContext(ctx,
		`INSERT INTO genre (name) VALUES ($1) RETURNING genre_id`,
		&r.Name).Scan(&r.GenreID)
	if err != nil {
		return errors.Wrap(err, "failed to insert genre")
	}
	return nil
}

// GetGenreByPkContext select the Genre from the database.
func GetGenreByPkContext(ctx context.Context, db Queryer, pk0 int) (*Genre, error) {
	var r Genre
	err := db.QueryRowContext(ctx,
		`SELECT genre_id, name FROM genre WHERE genre_id = $1`,
		pk0).Scan(&r.GenreID, &r.Name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select genre")
	}
	return &r, nil
}

// Hall represents public.hall
type Hall struct {
	HallID       int    // hall_id
	Title        string // title
	MovieHouseID int    // movie_house_id
}

// Create inserts the Hall to the database.
func (r *Hall) Create(db Queryer) error {
	return r.CreateContext(context.Background(), db)
}

// GetHallByPk select the Hall from the database.
func GetHallByPk(db Queryer, pk0 int) (*Hall, error) {
	return GetHallByPkContext(context.Background(), db, pk0)
}

// CreateContext inserts the Hall to the database.
func (r *Hall) CreateContext(ctx context.Context, db Queryer) error {
	err := db.QueryRowContext(ctx,
		`INSERT INTO hall (title, movie_house_id) VALUES ($1, $2) RETURNING hall_id`,
		&r.Title, &r.MovieHouseID).Scan(&r.HallID)
	if err != nil {
		return errors.Wrap(err, "failed to insert hall")
	}
	return nil
}

// GetHallByPkContext select the Hall from the database.
func GetHallByPkContext(ctx context.Context, db Queryer, pk0 int) (*Hall, error) {
	var r Hall
	err := db.QueryRowContext(ctx,
		`SELECT hall_id, title, movie_house_id FROM hall WHERE hall_id = $1`,
		pk0).Scan(&r.HallID, &r.Title, &r.MovieHouseID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select hall")
	}
	return &r, nil
}

// MovieHouse represents public.movie_house
type MovieHouse struct {
	MovieHouseID int    // movie_house_id
	Title        string // title
	Address      string // address
}

// Create inserts the MovieHouse to the database.
func (r *MovieHouse) Create(db Queryer) error {
	return r.CreateContext(context.Background(), db)
}

// GetMovieHouseByPk select the MovieHouse from the database.
func GetMovieHouseByPk(db Queryer, pk0 int) (*MovieHouse, error) {
	return GetMovieHouseByPkContext(context.Background(), db, pk0)
}

// CreateContext inserts the MovieHouse to the database.
func (r *MovieHouse) CreateContext(ctx context.Context, db Queryer) error {
	err := db.QueryRowContext(ctx,
		`INSERT INTO movie_house (title, address) VALUES ($1, $2) RETURNING movie_house_id`,
		&r.Title, &r.Address).Scan(&r.MovieHouseID)
	if err != nil {
		return errors.Wrap(err, "failed to insert movie_house")
	}
	return nil
}

// GetMovieHouseByPkContext select the MovieHouse from the database.
func GetMovieHouseByPkContext(ctx context.Context, db Queryer, pk0 int) (*MovieHouse, error) {
	var r MovieHouse
	err := db.QueryRowContext(ctx,
		`SELECT movie_house_id, title, address FROM movie_house WHERE movie_house_id = $1`,
		pk0).Scan(&r.MovieHouseID, &r.Title, &r.Address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select movie_house")
	}
	return &r, nil
}

// Person represents public.person
type Person struct {
	PersonID int    // person_id
	FullName string // full_name
}

// Create inserts the Person to the database.
func (r *Person) Create(db Queryer) error {
	return r.CreateContext(context.Background(), db)
}

// GetPersonByPk select the Person from the database.
func GetPersonByPk(db Queryer, pk0 int) (*Person, error) {
	return GetPersonByPkContext(context.Background(), db, pk0)
}

// CreateContext inserts the Person to the database.
func (r *Person) CreateContext(ctx context.Context, db Queryer) error {
	err := db.QueryRowContext(ctx,
		`INSERT INTO person (full_name) VALUES ($1) RETURNING person_id`,
		&r.FullName).Scan(&r.PersonID)
	if err != nil {
		return errors.Wrap(err, "failed to insert person")
	}
	return nil
}

// GetPersonByPkContext select the Person from the database.
func GetPersonByPkContext(ctx context.Context, db Queryer, pk0 int) (*Person, error) {
	var r Person
	err := db.QueryRowContext(ctx,
		`SELECT person_id, full_name FROM person WHERE person_id = $1`,
		pk0).Scan(&r.PersonID, &r.FullName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select person")
	}
	return &r, nil
}

// Place represents public.place
type Place struct {
	PlaceID    int // place_id
	PlaceInRow int // place_in_row
	Row        int // row
	HallID     int // hall_id
}

// Create inserts the Place to the database.
func (r *Place) Create(db Queryer) error {
	return r.CreateContext(context.Background(), db)
}

// GetPlaceByPk select the Place from the database.
func GetPlaceByPk(db Queryer, pk0 int) (*Place, error) {
	return GetPlaceByPkContext(context.Background(), db, pk0)
}

// CreateContext inserts the Place to the database.
func (r *Place) CreateContext(ctx context.Context, db Queryer) error {
	err := db.QueryRowContext(ctx,
		`INSERT INTO place (place_in_row, row, hall_id) VALUES ($1, $2, $3) RETURNING place_id`,
		&r.PlaceInRow, &r.Row, &r.HallID).Scan(&r.PlaceID)
	if err != nil {
		return errors.Wrap(err, "failed to insert place")
	}
	return nil
}

// GetPlaceByPkContext select the Place from the database.
func GetPlaceByPkContext(ctx context.Context, db Queryer, pk0 int) (*Place, error) {
	var r Place
	err := db.QueryRowContext(ctx,
		`SELECT place_id, place_in_row, row, hall_id FROM place WHERE place_id = $1`,
		pk0).Scan(&r.PlaceID, &r.PlaceInRow, &r.Row, &r.HallID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select place")
	}
	return &r, nil
}

// Session represents public.session
type Session struct {
	SessionID int       // session_id
	Date      time.Time // date
	Price     int       // price
	FormatID  int       // format_id
	FilmID    int       // film_id
	HallID    int       // hall_id
}

// Create inserts the Session to the database.
func (r *Session) Create(db Queryer) error {
	return r.CreateContext(context.Background(), db)
}

// GetSessionByPk select the Session from the database.
func GetSessionByPk(db Queryer, pk0 int) (*Session, error) {
	return GetSessionByPkContext(context.Background(), db, pk0)
}

// CreateContext inserts the Session to the database.
func (r *Session) CreateContext(ctx context.Context, db Queryer) error {
	err := db.QueryRowContext(ctx,
		`INSERT INTO session (date, price, format_id, film_id, hall_id) VALUES ($1, $2, $3, $4, $5) RETURNING session_id`,
		&r.Date, &r.Price, &r.FormatID, &r.FilmID, &r.HallID).Scan(&r.SessionID)
	if err != nil {
		return errors.Wrap(err, "failed to insert session")
	}
	return nil
}

// GetSessionByPkContext select the Session from the database.
func GetSessionByPkContext(ctx context.Context, db Queryer, pk0 int) (*Session, error) {
	var r Session
	err := db.QueryRowContext(ctx,
		`SELECT session_id, date, price, format_id, film_id, hall_id FROM session WHERE session_id = $1`,
		pk0).Scan(&r.SessionID, &r.Date, &r.Price, &r.FormatID, &r.FilmID, &r.HallID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select session")
	}
	return &r, nil
}

// Ticket represents public.ticket
type Ticket struct {
	TicketID   int  // ticket_id
	Paid       bool // paid
	SessionID  int  // session_id
	PlaceID    int  // place_id
	ConsumerID int  // consumer_id
}

// Create inserts the Ticket to the database.
func (r *Ticket) Create(db Queryer) error {
	return r.CreateContext(context.Background(), db)
}

// GetTicketByPk select the Ticket from the database.
func GetTicketByPk(db Queryer, pk0 int) (*Ticket, error) {
	return GetTicketByPkContext(context.Background(), db, pk0)
}

// CreateContext inserts the Ticket to the database.
func (r *Ticket) CreateContext(ctx context.Context, db Queryer) error {
	err := db.QueryRowContext(ctx,
		`INSERT INTO ticket (paid, session_id, place_id, consumer_id) VALUES ($1, $2, $3, $4) RETURNING ticket_id`,
		&r.Paid, &r.SessionID, &r.PlaceID, &r.ConsumerID).Scan(&r.TicketID)
	if err != nil {
		return errors.Wrap(err, "failed to insert ticket")
	}
	return nil
}

// GetTicketByPkContext select the Ticket from the database.
func GetTicketByPkContext(ctx context.Context, db Queryer, pk0 int) (*Ticket, error) {
	var r Ticket
	err := db.QueryRowContext(ctx,
		`SELECT ticket_id, paid, session_id, place_id, consumer_id FROM ticket WHERE ticket_id = $1`,
		pk0).Scan(&r.TicketID, &r.Paid, &r.SessionID, &r.PlaceID, &r.ConsumerID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select ticket")
	}
	return &r, nil
}

// Queryer database/sql compatible query interface
type Queryer interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}
