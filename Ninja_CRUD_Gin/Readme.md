# Ninja Weapons API

This project is a simple API built using the Gin framework in Go. It allows users to manage a collection of ninja weapons by performing CRUD (Create, Read, Update, Delete) operations.

# What is Gin?

Gin is a web framework written in Go (Golang). It provides a minimalistic yet powerful API for building web applications and microservices. Gin is known for its high performance and low memory footprint, making it an excellent choice for building efficient and scalable web applications.

## Key Features of Gin

- **Fast routing**: Gin uses a radix tree-based routing engine that allows for fast request routing.
- **Middleware support**: Gin supports middleware, which are functions that can be executed before or after the main handler, enabling functionalities like logging, authentication, and more.
- **Error handling**: Gin provides built-in mechanisms for handling errors, making it easier to manage errors in web applications.
- **JSON/XML rendering**: Gin includes built-in support for rendering JSON and XML responses, simplifying the process of serializing data for API endpoints.
- **Validation**: Gin offers support for request validation, allowing developers to easily validate incoming request data.

## Where is Gin Used?

Gin is commonly used in various scenarios, including:

- **Building RESTful APIs**: Gin's simplicity and performance make it well-suited for developing APIs that serve JSON or XML data over HTTP.
- **Microservices architecture**: Due to its lightweight nature, Gin is often used in microservices architectures where each service needs to handle a high volume of requests efficiently.
- **Single-page applications (SPAs)**: Gin can serve as the backend for SPAs built with frontend frameworks like React, Angular, or Vue.js, providing the necessary APIs and serving static files.
- **Prototyping and rapid development**: Gin's simplicity and ease of use make it an excellent choice for prototyping new ideas or rapidly developing web applications.


## Endpoints

### GET /ping

- **Description**: Check if the server is running.
- **Response**: Returns a JSON object with a "message" field set to "pong".

### GET /weapon

- **Description**: Retrieve information about a specific ninja weapon.
- **Parameters**:
  - Query Parameter: type (string) - The type of the weapon.
- **Response**:
  - 200 OK: Returns a JSON object with the details of the requested weapon.
  - 404 Not Found: If the requested weapon type does not exist.

### POST /weapon

- **Description**: Add a new ninja weapon to the collection.
- **Parameters**:
  - Query Parameter: type (string) - The type of the weapon.
  - Query Parameter: name (string) - The name of the weapon.
- **Response**:
  - 201 Created: If the weapon is added successfully.
  - 400 Bad Request: If either the weapon type or name is missing.
  - 409 Conflict: If the weapon already exists.

### DELETE /weapon

- **Description**: Remove a ninja weapon from the collection.
- **Parameters**:
  - Query Parameter: type (string) - The type of the weapon to delete.
- **Response**:
  - 200 OK: If the weapon is deleted successfully.
  - 404 Not Found: If the requested weapon type does not exist.

## Usage

1. Clone the repository.
2. Install dependencies using `go mod tidy`.
3. Run the main file `main.go`.
4. The server will start running on port 8080.
