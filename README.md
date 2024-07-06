# go-gin-books-api
## Description

The go-gin-books-api project is a simple RESTful API for managing a collection of books, built using the Go programming language and the Gin web framework. This project demonstrates the basic CRUD (Create, Read, Update, Delete) operations, providing a foundation for building more complex APIs and web applications.
Features

- Create: Add new books to the collection with details like title and author.
- Read: Retrieve the list of all books or a single book by its ID.
- Update: Modify the details of an existing book.
- Delete: Remove a book from the collection.

Technologies Used

- Go: A statically typed, compiled programming language designed for simplicity and efficiency.
- Gin: A fast, lightweight web framework for Go, known for its performance and simplicity.

## Getting Started
### Prerequisites

- Go (version 1.18 or higher)

### Installation

- Clone the repository:

```sh
git clone https://github.com/ombima56/go-gin-books-api.git
cd go-gin-books-api
```

### Install dependencies:

```sh
go mod tidy
```

Run the application:

```sh
go run `main.go`
```
The API will be running at `http://localhost:8080`.
## API Endpoints

- Get all books: GET /books
- Get a book by ID: GET /books/:id
- Create a new book: POST /books (with JSON body)
- Update a book: PUT /books/:id (with JSON body)
- Delete a book: DELETE /books/:id

### Examples
#### Get all books

```sh
curl -X GET http://localhost:8080/books
```
#### Get a book by ID

```sh
curl -X GET http://localhost:8080/books/1
```
#### Create a new book

```sh
curl -X POST http://localhost:8080/books -H "Content-Type: application/json" -d '{"title": "How to stop worrying and start living", "author": "Dake Carnegie"}'
```
#### Update a book

```sh
curl -X PUT http://localhost:8080/books/1 -H "Content-Type: application/json" -d '{"title": "The monk who sold his ferrari", "author": "Robin Sharma"}'
```
#### Delete a book

``` sh
curl -X DELETE http://localhost:8080/books/1
```
## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.