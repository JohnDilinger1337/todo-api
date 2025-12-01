# 1. Use official Go image
FROM golang:1.25.4-alpine AS build 

# 2. Enable CGO (needed for go-sqlite3)
ENV CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

# ensure go installed binaries are on PATH (where `go install` places them)
ENV PATH="/go/bin:${PATH}"

# 3. Install build tools for CGO and update packages
RUN apk update && apk upgrade && apk add --no-cache build-base git && \
    rm -rf /var/cache/apk/*

# Install `air` for live-reload during development
# pin a specific version to avoid surprises
RUN go install github.com/air-verse/air@latest

# 4. Set working directory
WORKDIR /app

# 5. Copy go.mod & go.sum first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# 6. Copy source code
COPY . .

# 7. Build the app
RUN go build -o server main.go

# 8. Expose port
EXPOSE 8080

# 9. Run the app
CMD ["./server"]

# Development stage with live-reload
FROM build AS development
CMD ["air"]
# Note: In development mode, the application will run with live-reload using `air`.
# Make sure to mount your local source code into the container for live-reload to work.
# You can do this using docker-compose or by using the `-v` flag with `docker run`.
# Example:
# docker run -p 8080:8080 -v $(pwd):/app your_image_name
# This will mount the current directory into /app in the container.
# The application will restart automatically on code changes.
# Also, ensure that your `air.toml` configuration file is set up correctly for your project structure.

