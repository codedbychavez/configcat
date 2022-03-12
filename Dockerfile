# ------ DEVELOPMENT BLOCK ------ #
FROM golang:1.17.6 as development

# Add a work directory
WORKDIR /app

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy app files
COPY . .

# Install Reflex for development -> Enables live reload
RUN go install github.com/cespare/reflex@latest

# Start app
CMD reflex -g '*.go' go run main.go --start-service