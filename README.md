# Packs Calculator

This project is a web application that calculates the number of packs needed for a given number using a specified set of pack sizes. It uses Go for the backend and Bootstrap for the frontend.

## Features

- Input a number and a list of pack sizes.
- Calculate the number of each pack size needed to fulfill the input number.
- Display the result in a user-friendly format.

## Technologies Used

- Go
- Echo framework
- Bootstrap
- jQuery
- Docker

## Getting Started

### Prerequisites

- Go 1.20 or later
- Docker

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/pr1sonmike/golang-packing-challenge.git
    cd golang-packing-challenge
    ```

2. Install Go dependencies:
    ```sh
    go mod tidy
    ```

3. Run the application:
    ```sh
    go run cmd/main.go
    ```

4. Open your browser and navigate to `http://localhost:8080`.

### Using Docker

1. Build the Docker image:
    ```sh
    docker build -t golang-packing-challenge .
    ```

2. Run the Docker container:
    ```sh
    docker run -p 8080:8080 golang-packing-challenge
    ```

## Project Structure

- `main.go`: Entry point of the application.
- `internal/handlers/packs_handler.go`: Contains the handler for calculating packs.
- `internal/service/packs_service.go`: Contains the business logic for fulfilling items.
- `templates/index.html`: Frontend template with Bootstrap and jQuery.
- `Dockerfile`: Docker configuration for building the application image.
- `.dockerignore`: Specifies files to ignore when building the Docker image.

## Usage

1. Enter a number in the "Number" field.
2. Enter pack sizes separated by commas in the "Packs" field.
3. Click "Calculate" to see the result.

## Running Tests

To run the tests, use the following command:
```sh
go test ./...