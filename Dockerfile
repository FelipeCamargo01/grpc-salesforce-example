FROM golang:1.23.0-bookworm

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /golang-rpc-server

# Optional: Expose port 8080
EXPOSE 8080

RUN apt-get update && apt-get install -y protobuf-compiler
RUN export PATH="$PATH:$(go env GOPATH)/bin"

CMD ["/golang-rpc-server"]