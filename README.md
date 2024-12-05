# Go Project Exercise

This project is an exercise for working with the Go programming language. The goal of this project is to familiarize yourself with the GET method and to use the Gin and GORM frameworks.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/dl/)

## Installation and Setup

1. Clone this repository:

   ```bash
   git clone https://github.com/AliGhanizade/sample_user_pas.git
   ```

2. Navigate to the project directory:

   ```bash
   cd repository
   ```

3. Install the dependencies:

   ```bash
   go mod tidy
   ```

4. Run the application:

   ```bash
   go run main.go
   ```

## Usage

After running the application, you can access the following URL using a browser or tools like Postman:

```
http://localhost:8080/
```

This endpoint connects to the GET method and returns the relevant data.

## Project Structure

- `main.go`: Entry point of the application.
- `routes.go`: Defines the routes and methods.
- `config/`: Contains GORM models for database interaction.
- `controllers/`: Contains controllers for handling requests.


