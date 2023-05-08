CREATE TABLE IF NOT EXISTS Country
(
    country_id  SERIAL NOT NULL,
    name VARCHAR(100) NOT NULL UNIQUE,
    PRIMARY KEY (country_id)
);

CREATE TABLE IF NOT EXISTS Film
(
    film_id  SERIAL NOT NULL,
    title VARCHAR(300) NOT NULL,
    description TEXT NOT NULL,
    duration INTEGER NOT NULL,
    release_date DATE NOT NULL,
    age_rating VARCHAR(20) NOT NULL,
    producer_id INTEGER NOT NULL,
    PRIMARY KEY (film_id)
);

CREATE TABLE IF NOT EXISTS Format
(
    format_id  SERIAL NOT NULL,
    name VARCHAR(20) NOT NULL UNIQUE,
    PRIMARY KEY (format_id)
);

CREATE TABLE IF NOT EXISTS Genre
(
    genre_id  SERIAL NOT NULL,
    name VARCHAR(200) NOT NULL UNIQUE,
    PRIMARY KEY (genre_id)
);

CREATE TABLE IF NOT EXISTS Hall
(
    hall_id  SERIAL NOT NULL,
    title VARCHAR(100) NOT NULL,
    movie_house_id INTEGER NOT NULL,
    PRIMARY KEY (hall_id)
);

CREATE TABLE IF NOT EXISTS Movie_house
(
    movie_house_id  SERIAL NOT NULL,
    title VARCHAR(100) NOT NULL,
    address VARCHAR(150) NOT NULL,
    PRIMARY KEY (movie_house_id)
);

CREATE TABLE IF NOT EXISTS Person
(
    person_id  SERIAL NOT NULL,
    full_name VARCHAR(200) NOT NULL,
    PRIMARY KEY (person_id)
);

CREATE TABLE IF NOT EXISTS Place
(
    place_id  SERIAL NOT NULL,
    place_in_row INTEGER NOT NULL,
    row INTEGER NOT NULL,
    hall_id INTEGER NOT NULL,
    PRIMARY KEY (place_id)
);

CREATE TABLE IF NOT EXISTS Session
(
    session_id  SERIAL NOT NULL,
    date TIMESTAMP NOT NULL,
    price INTEGER NOT NULL,
    format_id INTEGER NOT NULL,
    film_id INTEGER NOT NULL,
    hall_id INTEGER NOT NULL,
    PRIMARY KEY (session_id)
);

CREATE TABLE IF NOT EXISTS Ticket
(
    ticket_id  SERIAL NOT NULL,
    paid boolean NOT NULL,
    session_id INTEGER NOT NULL,
    place_id INTEGER NOT NULL,
    consumer_id INTEGER NOT NULL,
    PRIMARY KEY (ticket_id)
);

CREATE TABLE IF NOT EXISTS Consumer
(
    consumer_id  SERIAL NOT NULL,
    login VARCHAR(30) NOT NULL UNIQUE,
    password VARCHAR(60) NOT NULL,
    PRIMARY KEY (consumer_id)
);

CREATE TABLE IF NOT EXISTS Film_actor
(
    person_id INTEGER NOT NULL,
    film_id INTEGER NOT NULL,
    PRIMARY KEY (person_id, film_id)
);

CREATE TABLE IF NOT EXISTS Film_genre
(
    genre_id INTEGER NOT NULL,
    film_id INTEGER NOT NULL,
    PRIMARY KEY (genre_id, film_id)
);

CREATE TABLE IF NOT EXISTS Film_country
(
    country_id INTEGER NOT NULL,
    film_id INTEGER NOT NULL,
    PRIMARY KEY (country_id, film_id)
);


ALTER TABLE Film
    ADD FOREIGN KEY  (producer_id) REFERENCES Person (person_id);

ALTER TABLE Hall
    ADD FOREIGN KEY (movie_house_id) REFERENCES Movie_house (movie_house_id) ON DELETE CASCADE;

ALTER TABLE Place
    ADD FOREIGN KEY  (hall_id) REFERENCES Hall (hall_id) ON DELETE CASCADE;

ALTER TABLE Session
    ADD FOREIGN KEY  (format_id) REFERENCES Format (format_id),
    ADD FOREIGN KEY  (film_id) REFERENCES Film (film_id) ON DELETE CASCADE,
    ADD FOREIGN KEY  (hall_id) REFERENCES Hall (hall_id) ON DELETE CASCADE;

ALTER TABLE Ticket
    ADD FOREIGN KEY  (session_id) REFERENCES Session (session_id),
    ADD FOREIGN KEY  (place_id) REFERENCES Place (place_id),
    ADD FOREIGN KEY (consumer_id) REFERENCES Consumer (consumer_id) ;

ALTER TABLE Film_actor
    ADD FOREIGN KEY (film_id) REFERENCES Film (film_id) ON DELETE CASCADE,
    ADD FOREIGN KEY (person_id) REFERENCES Person (person_id)  ;

ALTER TABLE Film_genre
    ADD FOREIGN KEY (film_id) REFERENCES Film (film_id) ON DELETE CASCADE,
    ADD FOREIGN KEY (genre_id) REFERENCES Genre (genre_id)  ;

ALTER TABLE Film_country
    ADD FOREIGN KEY (film_id) REFERENCES Film (film_id) ON DELETE CASCADE,
    ADD FOREIGN KEY (country_id) REFERENCES Country (country_id)  ;