# File Converter

This is a Go project that provides APIs for converting files.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.22.2 or later)
- Docker

### Installing

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/yourrepository.git
    ```
2. Navigate to the project directory:
    ```sh
    cd yourrepository
    ```
3. Download the Go dependencies:
    ```sh
    go mod download
    ```

## Running the Application

You can run the application using the `go run` command:

```sh
go run main.go
```
The application will start a server on port 8080.

## Docker
This project includes a Dockerfile for building a Docker image of the application. You can build the image using the following command:
```sh
docker build -t yourimagename .
```

And then run the image with:
```sh
docker run -p 8080:8080 yourimagename
```

This will start the application inside a Docker container and map port 8080 of the container to port 8080 of your host machine.

## Endpoints:

* `POST /convert`: Converts a file.
* `POST /convert-img`: Converts an image file.
* `GET /metrics`: Returns Prometheus metrics.

