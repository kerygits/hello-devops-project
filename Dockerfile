FROM golang:1.21.1-alpine3.18
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN go build -o main
EXPOSE 11130/tcp
CMD [ "/app/main" ]
