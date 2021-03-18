FROM golang:latest

WORKDIR /
RUN mkdir /app
COPY . /app/

WORKDIR /app/src

ENV GOPATH /go

RUN go get gorm.io/gorm
RUN go get gorm.io/driver/postgres


RUN go get github.com/oxequa/realize
RUN go get github.com/gin-gonic/gin

RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/gin-contrib/cors
