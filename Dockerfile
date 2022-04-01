FROM golang:alpine

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./code/go.mod ./code/go.sum ./
RUN go mod download && go mod verify

COPY ./code .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]