FROM golang:latest

WORKDIR /
RUN mkdir /app
COPY . /app/

WORKDIR /app/src

ENV GOPATH /go
ENV GO111MODULE off

RUN go get gorm.io/gorm
RUN go get gorm.io/driver/postgres

# RUN go get github.com/urfave/cli/v2
# RUN go get gopkg.in/urfave/cli.v2@master
# RUN replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.2.0
# RUN GO111MODULE=off go get github.com/oxequa/realize
RUN go get -u github.com/cosmtrek/air
RUN go get github.com/gin-gonic/gin

RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/gin-contrib/cors

# RUN go mod init github.com/onyanko-pon/morning_compe_server