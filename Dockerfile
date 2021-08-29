FROM golang:1.16-alpine

RUN go get -u github.com/cosmtrek/air

RUN mkdir /todo-server

WORKDIR /todo-server
COPY . .

EXPOSE 3000