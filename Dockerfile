# 1. Use official Go image
FROM golang:1.21.6-alpine

# 2. Enable CGO (needed for go-sqlite3)
ENV CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

# 3. Install build tools for CGO and update packages
RUN apk update && apk upgrade && apk add --no-cache build-base && \
    rm -rf /var/cache/apk/*

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
