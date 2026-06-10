# Book Management API (Gin Gonic)

A lightweight, memory-backed RESTful API built in Go using the **Gin Web Framework**. This project provides standard CRUD operations to manage a collection of books, allowing you to create, read, update (checkout/return), and delete records seamlessly.

## Features

- **Full CRUD Support**: Add, view, find, and delete books instantly.
- **Transactional Simulation**: Simulate library transactions through explicit checkout and return mechanics.
- **In-Memory Data Store**: Clean, lightweight architecture using safe in-memory data structures.
- **Robust Error Handling**: Handles missing query parameters, invalid routes, and structured JSON error responses.

---

## Tech Stack

- **Language**: Go (Golang)
- **Framework**: [Gin Web Framework](https://github.com/gin-gonic/gin)
- **Data Format**: JSON

---

## Project Structure

- `main.go`: Contains the foundational route handlers, model declarations (`book` struct), application configuration, and initialization routines.
- `body.json`: Sample payload structure representing a model instance for API integration and requests.

---

## 🔌 API Endpoints

### 1. Retrieve All Books
* **Method**: `GET`
* **Route**: `/books`
* **Success Response**: `200 OK`

### 2. Retrieve a Specific Book
* **Method**: `GET`
* **Route**: `/books/:id`
* **Success Response**: `200 OK`
* **Error Response**: `404 Not Found` (If the requested ID does not exist)

### 3. Insert a New Book
* **Method**: `POST`
* **Route**: `/books`
* **Payload Format**: JSON (Bindable to the core Struct model)
* **Success Response**: `201 Created`

### 4. Checkout a Book (Decrement Stock)
* **Method**: `PATCH`
* **Route**: `/checkout?id=<book_id>`
* **Success Response**: `200 OK` (Decrements `quantity` field by 1)
* **Error Responses**:
    - `400 Bad Request` (Missing `id` query parameter or stock level is `0`)
    - `404 Not Found` (Invalid book ID)

### 5. Return a Book (Increment Stock)
* **Method**: `PATCH`
* **Route**: `/return?id=<book_id>`
* **Success Response**: `200 OK` (Increments `quantity` field by 1)
* **Error Responses**:
    - `400 Bad Request` (Missing `id` query parameter)
    - `404 Not Found` (Invalid book ID)

### 6. Remove a Book
* **Method**: `DELETE`
* **Route**: `/books/:id`
* **Success Response**: `200 OK`
* **Error Response**: `404 Not Found`

---

## How to Run the Project

Follow these steps to set up and run the API application locally on your machine.

### Prerequisites
Make sure you have **Go** installed on your system (version 1.16+ is highly recommended). You can verify your installation by running:
```bash
go version
```


## Part 1: Initialize the Module

If you haven't initialized a Go module in your project directory yet, open your terminal and run:

```bash
go mod init book-api
```

> **Note:** If a `go.mod` file already exists in your directory, you can safely skip this step.

---

## Part 2: Install Dependencies

Download and install the Gin Web Framework package required for this project:

```bash
go get github.com/gin-gonic/gin
```

---

## Part 3: Start the Server

Compile and run the primary application file:

```bash
go run main.go
```

Once executed, the development server will spin up and dynamically bind to your local port:

```text
[GIN-debug] Listening and serving HTTP on localhost:8080
```

---

# Testing with cURL

With the server running on `localhost:8080`, open a separate terminal window to interact with the endpoints.

## Fetch All Books

```bash
curl http://localhost:8080/books
```

## Add a Book (Using the Sample Payload File)

```bash
curl -X POST http://localhost:8080/books \
-H "Content-Type: application/json" \
-d @body.json
```

## Add a Book (Explicit JSON String)

```bash
curl -X POST http://localhost:8080/books \
-H "Content-Type: application/json" \
-d '{"id": "5", "title": "Atomic Habits", "author": "James Clear", "quantity": 1000}'
```

## Checkout a Book

```bash
curl -X PATCH "http://localhost:8080/checkout?id=2"
```

## Delete a Book

```bash
curl -X DELETE http://localhost:8080/books/1
```
