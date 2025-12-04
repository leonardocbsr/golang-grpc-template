FROM golang:1.24-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./gotemplate .

EXPOSE 8080
ENTRYPOINT ["./gotemplate"]
