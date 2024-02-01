# echo-api

This is a simple Go application.

## Structure

The application has the following structure:

- `app/go.mod`: The Go module file.
- `app/main.go`: The main entry point of the application.
- `app/Dockerfile`: The Dockerfile for creating a Docker image of the
  application.

## Setup

To set up the application, follow these steps:

1. Ensure that you have Go and Docker installed on your machine.
2. Navigate to the `app` directory.
3. Run `go mod download` to download the necessary Go dependencies.
4. Run `go build` to build the application.
5. Run `docker build -t my-go-app .` to build the Docker image.

## Usage

To run the application, use the following command:

```
docker run -p 8080:8080 my-go-app
```

This will start the application and make it available at
`http://localhost:8080`.

```
Please adjust the setup and usage instructions as necessary based on the specifics of your application.
```
