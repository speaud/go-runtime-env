FROM golang:1.16

WORKDIR /api

COPY ./ /api

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main