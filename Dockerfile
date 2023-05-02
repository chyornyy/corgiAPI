FROM golang:1.20-alpine
RUN apk add --no-cache inotify-tools
RUN mkdir /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
CMD ["./main"]
ENV GIN_MODE=debug
EXPOSE 8080
