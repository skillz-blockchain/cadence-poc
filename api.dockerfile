FROM golang:1.16

WORKDIR /api

COPY . .
RUN go build -o /bin/cadence-poc-api cmd/api.go

CMD /bin/cadence-poc-api