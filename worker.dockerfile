FROM golang:1.16

FROM golang:1.16

WORKDIR /worker

COPY . .
RUN go build -o /bin/cadence-poc-worker cmd/worker.go

CMD /bin/cadence-poc-worker