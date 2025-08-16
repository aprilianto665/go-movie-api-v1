CREATE TABLE IF NOT EXISTS movie (
  id           INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  title        TEXT NOT NULL,
  image        TEXT,
  category     TEXT,
  genre        TEXT[] NOT NULL DEFAULT '{}',
  duration     INTEGER NOT NULL CHECK (duration > 0),
  release_year INTEGER NOT NULL CHECK (release_year >= 1888),
  age_rating   INTEGER NOT NULL CHECK (age_rating >= 0),
  synopsis     TEXT,
  "cast"       TEXT[] NOT NULL DEFAULT '{}',
  creators     TEXT[] NOT NULL DEFAULT '{}',
  rating       DOUBLE PRECISION NOT NULL CHECK (rating >= 0 AND rating <= 10)
);

INSERT INTO movie (title, image, category, genre, duration, release_year, age_rating, synopsis, "cast", creators, rating)
VALUES
(
  'The Shawshank Redemption',
  'https://example.com/shawshank.jpg',
  'FILM',
  ARRAY['Drama', 'Crime']::text[],
  142,
  1994,
  15,
  'Two imprisoned men bond over a number of years, finding solace and eventual redemption through acts of common decency.',
  ARRAY['Tim Robbins', 'Morgan Freeman', 'Bob Gunton']::text[],
  ARRAY['Frank Darabont', 'Niki Marvin']::text[],
  9.3
),
(
  'The Godfather',
  'https://example.com/godfather.jpg',
  'FILM',
  ARRAY['Crime', 'Drama']::text[],
  175,
  1972,
  18,
  'The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.',
  ARRAY['Marlon Brando', 'Al Pacino', 'James Caan']::text[],
  ARRAY['Francis Ford Coppola', 'Albert S. Ruddy']::text[],
  9.2
),
(
  'The Dark Knight',
  'https://example.com/darkknight.jpg',
  'FILM',
  ARRAY['Action', 'Crime', 'Drama']::text[],
  152,
  2008,
  13,
  'When the menace known as the Joker wreaks havoc and chaos on the people of Gotham, Batman must accept one of the greatest psychological and physical tests.',
  ARRAY['Christian Bale', 'Heath Ledger', 'Aaron Eckhart']::text[],
  ARRAY['Christopher Nolan', 'Emma Thomas']::text[],
  9.0
);