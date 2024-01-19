# Simple Go Server with Fiber

This repository contains a simple Go server built with the Fiber web framework.

## Features

- Lightweight and fast HTTP server.
- Utilizes Fiber, an Express.js inspired web framework for Go.
- Quick and easy setup.

## Requirements

- Go (version X.X.X)
- Fiber (version X.X.X)

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/your-repo.git
   cd your-repo
   go mod download
   go run main.go
   ```

   The server will be running at http://localhost:3000.

## Usage

This is a simple movies server that provides a RESTful API for managing a list of movies. Below are the available endpoints and features:

### Get All Movies

**Endpoint:**
GET /api/movies

**Description:**
Get a list of all movies in JSON format.

### Get a Single Movie by ID

**Endpoint:**
GET /api/movies/:id

**Description:**
Get details of a single movie by providing its unique ID in the request parameter.

### Create a New Movie

**Endpoint:**
POST /api/movies

**Description:**
Create a new movie by sending a POST request with the movie details in the request body - eg: {"title": "Movie Name", "director": "Director Name"}.

### Delete a Movie by ID

**Endpoint:**
DELETE /api/movies/:id

**Description:**
Delete a movie by providing its unique ID in the request parameter.
