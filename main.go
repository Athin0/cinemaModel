package main

import (
	PgRepo "cinema/db"
	"context"
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"time"
)

var apikey = ""

func main() {
	/*imbdIds := readCsvFile("./movies_metadata.csv")
	f, err := os.Open("film.json")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}*/

	db, err := PgRepo.InitDB()
	ctx := context.Background()
	if err != nil {
		log.Fatal(err)
	}
	/*
		for j := 70; j < 101; j++ {
			for i := 1; i < 12; i++ {
				_, err := db.Client.ExecContext(ctx, "INSERT INTO hall(title,movie_house_id) VALUES ($1,$2)", "Best Place"+strconv.Itoa(i), j)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	*/
	/*
		for j := 21; j < 101; j++ {
			for row := 1; row < 4; row++ {
				for place := 1; place < 7; place++ {
					_, err := db.Client.ExecContext(ctx, "INSERT INTO place(hall_id,row,place_in_row) VALUES ($1,$2,$3)", j, row, place)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}*/
	/*
		for j := 0; j < 1000; j++ {
			date := time.Now().Add(time.Hour * time.Duration(rand.Intn(2760))) //time.Date(2023,rand.Intn(11)+1,rand.Intn(27)+1,)
			_, err := db.Client.ExecContext(ctx, "insert into session (date, price, format_id, film_id, hall_id) VALUES ($1,$2,$3,$4,$5)", date, (rand.Intn(100)+15)*10, rand.Intn(8)+1, rand.Intn(99)+1, rand.Intn(100)+2)
			if err != nil {
				log.Fatal(err)
			}

		}
	*/

	for j := 0; j < 100; j++ {
		date := time.Now().Add(time.Hour * time.Duration(rand.Intn(2760))) //time.Date(2023,rand.Intn(11)+1,rand.Intn(27)+1,)
		_, err := db.Client.ExecContext(ctx, "insert into ticket (date, price, format_id, film_id, hall_id) VALUES ($1,$2,$3,$4,$5)", date, (rand.Intn(100)+15)*10, rand.Intn(8)+1, rand.Intn(99)+1, rand.Intn(100)+2)
		if err != nil {
			log.Fatal(err)
		}

	}
}

/*
	client := omdb.NewOMBD(apikey)

	for _, id := range imbdIds {
		film := client.Read(id)
		bytes, err := json.Marshal(film)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write(bytes)
		if err != nil {
			log.Fatal(err)
		}


*/

/*
	err = db.AddFilm(ctx, film)
	if err != nil {
		log.Fatal(err)
	}*/

/*
file, err := os.Open("film.json")
	if err != nil {
		log.Fatal(err)
	}
	var film []models.FilmJson
	err = json.NewDecoder(file).Decode(&film)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(film)

*/

func readCsvFile(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	csvReader.Comma = ','
	var records []string
	field := 6
	_, err = csvReader.Read()
	begin := 50
	for i := 0; i < begin; i++ {
		csvReader.Read()
	}
	for i := begin; i < begin+100; i++ {
		record, err := csvReader.Read()
		records = append(records, record[field])
		if err != nil {
			log.Fatal("Unable to parse file as CSV for "+filePath, err)
		}
	}
	return records
}
