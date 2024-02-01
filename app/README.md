# echo-server

This application is a simple, lightweight microservice written in Go. It serves
as an echo API, responding to HTTP requests at its root URL with a JSON payload.
The payload includes the current timestamp and the IP address of the client
making the request.

The application is designed to be containerized using Docker, and it includes
security measures such as running as a non-root user within the container. This
makes it a suitable starting point for developing more complex Go-based
microservices.

## Structure

The application has the following structure:

- `app/go.mod`: The Go module file.
- `app/main.go`: The main entry point of the application.
- `app/Dockerfile`: The Dockerfile for creating a Docker image of the
  application.

## Build

To build the application, follow these steps:

1. Ensure that you have Docker installed on your machine.
2. Navigate to the `app` directory.
3. Run:

```sh
GOARCH=amd64  ## <-- example; default is arm64
IMG="yourdockerhuborg/echo-server"
docker build --tag "${IMG}" .
```

Set `IMG` to the correct value, including the tag. It is recommended to tag the
image with a uniquely identifiable value, such as the git commit hash.

## Run

To run the application, use the following command:

```sh
docker run --rm --publish 8080:8080 "${IMG}"
```

This will start the application and make it available at
`http://localhost:8080`.

Verify it is running with:

```sh
curl http://localhost:8080
```
