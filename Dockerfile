FROM golang:1.24.2-alpine AS go

ENV GOLANG_VERSION=1.24.2
ENV GOTOOLCHAIN=local
ENV GODIR=/opt/go
ENV GOAPP=$GODIR
# sqlite3 と go-sqlite3のためにこの環境変数が必要らしい（1.24時点では不要なのかも
ENV CGO_ENABLED=1

# アップデートとgitのインストール
RUN apk update && apk add git && apk add --no-cache ca-certificates 
# sqlite3 のインストール go-sqlite3のためにgccが必要
RUN apk add --no-cache gcc
RUN apk add musl-dev 
RUN apk add sqlite

RUN mkdir -p "$GOAPP/cmd" "$GOAPP/bin" 
RUN chmod -R 1777 "$GODIR"

WORKDIR "$GOAPP"
RUN cd "$GOAPP"

RUN go mod init go-sample-todo 
RUN go mod tidy

#RUN go get github.com/mattn/go-sqlite3
#RUN go get -u github.com/go-sql-driver/mysql
#RUN go get -u github.com/gin-gonic/gin 
#RUN go install github.com/air-verse/air@latest 

#RUN air init
#RUN cd "$GOAPP/cmd"; air &
