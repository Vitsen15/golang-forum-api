FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Andrey Vitsentovich <vitsentovich.andrey@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY ../../go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download -x

# Copy the source from the current directory to the Working Directory inside the container
COPY ../.. .

# Build the Go app
RUN go build -x -o main .

# Add docker-compose-wait tool -------------------
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

# Run the executable
CMD ["./main"]
