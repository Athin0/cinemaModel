package PgRepo

import (
	"cinema/models"
	"context"
	"database/sql"
	_ "encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type PostgresDB struct {
	Client *sql.DB
}

func NewPostgresDB(cfg Config) (*PostgresDB, error) {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("postgres connect error : (%v)", err)
	}
	fmt.Println(db)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresDB{Client: db}, nil
}

func InitDB() (*PostgresDB, error) {
	viper.AddConfigPath("../cinema/db") //what?

	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("error in reading config: %v", err)
		return nil, err
	}
	db, err := NewPostgresDB(Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("error creating db: %v \n", err)
		return nil, err
	}
	return db, nil
}

func (db *PostgresDB) GetCountry(name string) (int, error) {
	err := db.Client.QueryRow(
		"SELECT country_id, name FROM country WHERE name = $1",
		name,
	)
	if err.Err() != nil {
		log.Printf(err.Err().Error())
		return 0, err.Err()
	}
	var ans models.Country
	err2 := err.Scan(&ans.CountryID, &ans.Name)
	if err2 != nil {
		return 0, err2
	}
	return ans.CountryID, nil
}

func (db *PostgresDB) GetPerson(name string) (int, error) {
	err := db.Client.QueryRow(
		"SELECT person_id, full_name FROM person WHERE full_name = $1",
		name,
	)
	if err.Err() != nil {
		log.Printf(err.Err().Error())
		return 0, err.Err()
	}
	var ans models.Person
	err2 := err.Scan(&ans.PersonID, &ans.FullName)
	if err2 != nil {
		return 0, err2
	}
	return ans.PersonID, nil
}

func (db *PostgresDB) GetGenre(name string) (int, error) {
	err := db.Client.QueryRow(
		"SELECT genre_id, name FROM genre WHERE name = $1",
		name,
	)

	if err.Err() != nil {
		log.Printf(err.Err().Error())
		return 0, err.Err()
	}
	var ans models.Genre
	err2 := err.Scan(&ans.GenreID, &ans.Name)
	if err2 != nil {
		return 0, err2
	}
	return ans.GenreID, nil
}

func (db *PostgresDB) AddFilm(ctx context.Context, film models.FilmJson) error {
	tx, err := db.Client.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		tx.Rollback()
		return err
	}
	producer := &models.Person{FullName: film.Producer}
	producer.PersonID, err = db.GetPerson(producer.FullName)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			tx.Rollback()
			return err
		}
		err = producer.CreateContext(ctx, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	f := &models.Film{Title: film.Title, Description: film.Description, Duration: film.GetDuration(), ReleaseDate: film.GetDate(), AgeRating: film.AgeRating, ProducerID: producer.PersonID}
	err = f.CreateContext(ctx, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, a := range film.GetActors() {
		actor := &models.Person{FullName: a}
		actor.PersonID, err = db.GetPerson(actor.FullName)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				tx.Rollback()
				return err
			}
			err = actor.CreateContext(ctx, tx)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		filmActor := &models.FilmActor{FilmID: f.FilmID, PersonID: actor.PersonID}
		err = filmActor.CreateContext(ctx, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, a := range film.GetGenres() {
		genre := &models.Genre{Name: a}
		genre.GenreID, err = db.GetGenre(genre.Name)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				tx.Rollback()
				return err
			}
			err = genre.CreateContext(ctx, tx)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		filmGenre := &models.FilmGenre{FilmID: f.FilmID, GenreID: genre.GenreID}
		err = filmGenre.CreateContext(ctx, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, a := range film.GetCountries() {
		country := &models.Country{Name: a}
		country.CountryID, err = db.GetCountry(country.Name)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				tx.Rollback()
				return err
			}
			err = country.CreateContext(ctx, tx)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		filmCountry := &models.FilmCountry{FilmID: f.FilmID, CountryID: country.CountryID}
		err = filmCountry.CreateContext(ctx, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

/* Redundant (in models)

func (db *PostgresDB) AddCountry(name string) error {
	err := db.Client.QueryRow(
		"INSERT INTO country (name) VALUES ($1)",
		name,
	)
	if err.Err() != nil {
		log.Printf(err.Err().Error())
	}
	return nil

}
func (db *PostgresDB) AddPerson(name string) error {
	err := db.Client.QueryRow(
		"INSERT INTO person(full_name) VALUES ($1)",
		name,
	)
	if err.Err() != nil {
		log.Printf(err.Err().Error())
	}
	return nil
}
func (db *PostgresDB) AddGenre(name string) error {
	err := db.Client.QueryRow(
		"INSERT INTO genre(name) VALUES ($1)",
		name,
	)
	if err.Err() != nil {
		log.Printf(err.Err().Error())
	}
	return nil
}
func (db *PostgresDB) AddFilm(f models.Film) error {
	err := db.Client.QueryRow(
		"INSERT INTO film(T) VALUES ($1)",
		f.Title, f.Duration, f.Duration, f.ReleaseDate, f.AgeRating, f.CountryID, f.ProducerID,
	)
	if err.Err() != nil {
		log.Printf(err.Err().Error())
	}
	return nil
}*/
