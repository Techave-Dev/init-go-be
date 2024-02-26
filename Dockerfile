FROM golang:1.22.0-bullseye AS builder

WORKDIR /app

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.0

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.25.0

COPY go.mod go.sum ./

RUN go mod download

FROM golang:1.22.0-bullseye

WORKDIR /app

COPY --from=builder /go/bin/sqlc /bin/sqlc

COPY --from=builder /go/bin/migrate /bin/migrate

COPY --from=builder /go/pkg/mod /go/pkg/mod

COPY . .

CMD [ "make", "start" ]
