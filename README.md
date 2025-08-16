# Go Movie API v1

A RESTful API for managing movies built with Go and PostgreSQL using Repository Pattern.

## Features

- CRUD operations for movies
- PostgreSQL database with array support
- Repository pattern architecture
- Consistent JSON response format
- Input validation and error handling

## Tech Stack

- **Language**: Go 1.24+
- **Database**: PostgreSQL
- **Driver**: pgx/v5
- **Architecture**: Repository Pattern

## Project Structure

```
go-movie-api-v1/
├── controllers/        # HTTP handlers
├── models/            # Data models and response structs
├── repositories/      # Data access layer
├── routes/           # Route definitions
├── db/              # Database connection and schema
├── app.go           # Main application entry point
└── README.md        # This file
```

## Installation

### Prerequisites

- Go 1.24 or higher
- PostgreSQL 12+
- Git

### Setup

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd go-movie-api-v1
   ```

2. **Install dependencies**

   ```bash
   go mod tidy
   ```

3. **Setup database**

   ```bash
   # Create database
   createdb go_movie_v1

   # Run schema
   psql -U your_username -d go_movie_v1 -f db/schema.sql
   ```

4. **Configure database connection**

   Update the DSN in `app.go`:

   ```go
   dsn := "postgres://username:password@localhost:5432/go_movie_v1?sslmode=disable"
   ```

5. **Run the application**
   ```bash
   go run .
   ```

The server will start on `http://localhost:8080`

## API Endpoints

### Get All Movies

```http
GET /movies
```

**Response:**

```json
{
  "status": "success",
  "message": "Movies retrieved successfully",
  "data": [
    {
      "id": 1,
      "title": "Movie Title",
      "image": "https://example.com/image.jpg",
      "category": "FILM",
      "genre": ["Action", "Adventure"],
      "duration": 120,
      "releaseYear": 2023,
      "ageRating": 18,
      "synopsis": "Movie synopsis...",
      "cast": ["Actor 1", "Actor 2"],
      "creators": ["Director", "Producer"],
      "rating": 8.5
    }
  ]
}
```

### Get Movie by ID

```http
GET /movie/{id}
```

**Response:**

```json
{
  "status": "success",
  "message": "Movie retrieved successfully",
  "data": {
    "id": 1,
    "title": "Movie Title"
    // ... movie data
  }
}
```

### Create Movie

```http
POST /movie
Content-Type: application/json

{
  "title": "New Movie",
  "image": "https://example.com/image.jpg",
  "category": "FILM",
  "genre": ["Action", "Adventure"],
  "duration": 120,
  "releaseYear": 2023,
  "ageRating": 18,
  "synopsis": "Movie synopsis...",
  "cast": ["Actor 1", "Actor 2"],
  "creators": ["Director", "Producer"],
  "rating": 8.5
}
```

**Response:**

```json
{
  "status": "success",
  "message": "Movie created successfully",
  "data": {
    "id": 4,
    "title": "New Movie",
    "image": "https://example.com/image.jpg",
    "category": "FILM",
    "genre": ["Action", "Adventure"],
    "duration": 120,
    "releaseYear": 2023,
    "ageRating": 18,
    "synopsis": "Movie synopsis...",
    "cast": ["Actor 1", "Actor 2"],
    "creators": ["Director", "Producer"],
    "rating": 8.5
  }
}
```

### Update Movie

```http
PUT /movie/{id}
Content-Type: application/json

{
  "title": "Updated Movie Title",
  "image": "https://example.com/updated-image.jpg",
  "category": "SERIES",
  "genre": ["Drama", "Thriller"],
  "duration": 135,
  "releaseYear": 2024,
  "ageRating": 16,
  "synopsis": "Updated movie synopsis...",
  "cast": ["New Actor 1", "New Actor 2"],
  "creators": ["New Director", "New Producer"],
  "rating": 9.0
}
```

**Response:**

```json
{
  "status": "success",
  "message": "Movie updated successfully",
  "data": {
    "id": 1,
    "title": "Updated Movie Title",
    "image": "https://example.com/updated-image.jpg",
    "category": "SERIES",
    "genre": ["Drama", "Thriller"],
    "duration": 135,
    "releaseYear": 2024,
    "ageRating": 16,
    "synopsis": "Updated movie synopsis...",
    "cast": ["New Actor 1", "New Actor 2"],
    "creators": ["New Director", "New Producer"],
    "rating": 9.0
  }
}
```

### Delete Movie

```http
DELETE /movie/{id}
```

**Response:**

```json
{
  "status": "success",
  "message": "Movie deleted successfully",
  "data": {
    // deleted movie data
  }
}
```

## Database Schema

The `movie` table structure:

| Column       | Type             | Constraints                 |
| ------------ | ---------------- | --------------------------- |
| id           | INTEGER          | PRIMARY KEY, AUTO INCREMENT |
| title        | TEXT             | NOT NULL                    |
| image        | TEXT             | -                           |
| category     | TEXT             | -                           |
| genre        | TEXT[]           | NOT NULL, DEFAULT '{}'      |
| duration     | INTEGER          | NOT NULL, > 0               |
| release_year | INTEGER          | NOT NULL, >= 1888           |
| age_rating   | INTEGER          | NOT NULL, >= 0              |
| synopsis     | TEXT             | -                           |
| cast         | TEXT[]           | NOT NULL, DEFAULT '{}'      |
| creators     | TEXT[]           | NOT NULL, DEFAULT '{}'      |
| rating       | DOUBLE PRECISION | NOT NULL, 0-10              |

## Development

### Building

```bash
go build .
```

### Testing

```bash
go test ./...
```

### Code Structure

- **Repository Pattern**: Separates data access logic from business logic
- **Dependency Injection**: Controllers depend on repository interfaces
- **Consistent Responses**: All endpoints return standardized JSON format
- **Error Handling**: Proper HTTP status codes and error messages
