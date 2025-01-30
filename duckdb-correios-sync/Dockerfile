FROM golang:1.23 AS build

RUN apt-get update && apt-get install --no-install-recommends -y build-essential

WORKDIR /go/src
COPY . .

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags="no_duckdb_arrow" -ldflags="-linkmode external -extldflags '-static' -s -w" -o /go/bin/app

FROM scratch
WORKDIR /go
COPY --from=build /go/bin/app ./
CMD ["./app", "--postgres"]
