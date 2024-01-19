package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
}

type DB struct {
	movies []Movie
}

var db DB

func setMovies(db *DB) {
	db.movies = append(db.movies, Movie{ID: "1", Title: "The Shawshank Redemption", Director: "Frank Darabont"}, Movie{ID: "2", Title: "The Godfather", Director: "Francis Ford Coppola"}, Movie{ID: "3", Title: "The Dark Knight", Director: "Christopher Nolan"}, Movie{ID: "4", Title: "Pulp Fiction", Director: "Quentin Tarantino"}, Movie{ID: "5", Title: "Schindler's List", Director: "Steven Spielberg"}, Movie{ID: "6", Title: "Forrest Gump", Director: "Robert Zemeckis"}, Movie{ID: "7", Title: "Fight Club", Director: "David Fincher"}, Movie{ID: "8", Title: "The Matrix", Director: "The Wachowskis"}, Movie{ID: "9", Title: "Forrest Gump", Director: "Robert Zemeckis"}, Movie{ID: "10", Title: "Inception", Director: "Christopher Nolan"})
}

func (db *DB) getAllMovies(c *fiber.Ctx) error {

	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "Found All Movies", "data": db.movies, "length": len(db.movies)})

	return nil
}

func (db *DB) getMovieById(c *fiber.Ctx) error {

	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "ID cannot be empty"})
	} else {
		for _, movie := range db.movies {
			if movie.ID == id {
				c.Status(http.StatusOK).JSON(fiber.Map{"message": "Found the movie", "data": movie})
				return nil
			}
		}

		c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Could not find the movie with the given ID"})
	}

	return nil
}

func (db *DB) createMovie(c *fiber.Ctx) error {

	movie := Movie{}

	if err := c.BodyParser(&movie); err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Wrong information provided for the movie"})
		return nil
	}

	movie.ID = uuid.New().String()

	db.movies = append(db.movies, movie)

	c.Status(http.StatusOK).JSON(fiber.Map{"message": "Movie created Successfully", "data": movie})

	return nil
}

func (db *DB) deleteMovie(c *fiber.Ctx) error {

	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "ID cannot be empty"})
	} else {
		for index, movie := range db.movies {
			if movie.ID == id {
				db.movies = append(db.movies[:index], db.movies[index+1:]...)
				c.Status(http.StatusOK).JSON(fiber.Map{"message": "Movie Deleted", "data": movie})
				return nil
			}

		}
		c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Could not find the movie with the given ID"})
	}

	return nil
}

func setUpApiRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/movies", db.getAllMovies)
	api.Get("/movies/:id", db.getMovieById)
	api.Post("/movies", db.createMovie)
	api.Delete("/movies/:id", db.deleteMovie)
}

func main() {
	app := fiber.New()

	app.Get("/", func(context *fiber.Ctx) error {
		context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Welcome to Jeet's Go Server"})
		return nil
	})

	setMovies(&db)

	setUpApiRoutes(app)

	err := app.Listen(":8000")

	if err != nil {
		log.Fatal("Error: ", err)
	} else {
		fmt.Println("Server started on port 3000")
	}
}
