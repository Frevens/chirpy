# Chirpy

Chirpy is a RESTful API built with Go and PostgreSQL as part of the [Boot.dev Backend Developer Path](https://www.boot.dev/).

The project simulates the backend of a small social network where users can create, retrieve, and delete short messages called "chirps".

Although Chirpy was built as part of a guided curriculum, the project provided hands-on experience with building a complete backend service in Go, including authentication, database integration, HTTP handlers, middleware, and API design.

## Features

* User registration and login
* Password hashing
* JWT-based authentication
* Access and refresh tokens
* Token revocation
* User profile management
* Create and retrieve chirps
* Delete chirps with authorization checks
* Filter chirps by author
* Sort chirps by creation date
* PostgreSQL database persistence
* Database migrations
* SQL queries generated with `sqlc`
* HTTP middleware
* JSON API responses
* Error handling
* API endpoint testing

## Tech Stack

* **Go**
* **PostgreSQL**
* **sqlc**
* **Goose** for database migrations
* **JWT**
* **bcrypt/Argon2 password hashing**
* Standard Go `net/http` package

## API

The API provides endpoints for:

### Users

* `POST /api/users` — Create a user
* `PUT /api/users` — Update user information

### Authentication

* `POST /api/login` — Authenticate a user and obtain tokens
* `POST /api/refresh` — Refresh an access token
* `POST /api/revoke` — Revoke a refresh token

### Chirps

* `POST /api/chirps` — Create a chirp
* `GET /api/chirps` — Retrieve chirps
* `GET /api/chirps/{chirpID}` — Retrieve a specific chirp
* `DELETE /api/chirps/{chirpID}` — Delete a chirp

The chirp endpoint supports optional filtering and sorting:

```text
GET /api/chirps?author_id=<UUID>
GET /api/chirps?sort=asc
GET /api/chirps?sort=desc
```

`asc` is the default sort order.

## Database

Chirpy uses PostgreSQL for persistent storage.

Database schema changes are managed through migrations, while SQL queries are generated into type-safe Go code using `sqlc`.

The project uses a separation between:

* HTTP handlers
* API configuration
* database queries
* SQL migrations
* authentication logic
* utility functions

## Running Locally

### Requirements

* Go
* PostgreSQL
* `goose`
* `sqlc`

### Clone the repository

```bash
git clone https://github.com/Frevens/chirpy.git
cd chirpy
```

### Configure the database

Create a PostgreSQL database and configure the connection string used by the application.

The application expects the database URL to be provided through the environment.

### Run migrations

Apply the database migrations before starting the server.

### Start the server

```bash
go run .
```

The API will be available at:

```text
http://localhost:8080
```

## What I Learned

This project was part of my transition from learning individual programming concepts to building a complete backend application.

Some of the main concepts practiced were:

* Designing RESTful HTTP APIs
* Working with Go's `net/http`
* Structuring a Go backend project
* Connecting Go applications to PostgreSQL
* Writing and managing SQL migrations
* Using `sqlc` to generate database access code
* Implementing authentication and authorization
* Working with JWTs and refresh tokens
* Writing HTTP middleware
* Handling request validation and errors
* Working with query parameters
* Sorting and filtering API responses
* Testing an API against predefined requirements

## Project Context

This repository was created as part of the **Boot.dev Backend Developer Path**.

It is primarily a learning project rather than a production application. The goal was to gain practical experience building a backend service in Go and working with the technologies commonly used in backend development.

---

**Part of my journey learning backend development with Go.**

