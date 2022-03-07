FROM golang:1.17 AS build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o /bareska-interview

### Deploy built binary

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /bareska-interview /bareska-interview

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/bareska-interview"]