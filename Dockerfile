FROM golang:1.17 AS build

ARG LOG_FILE

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ${LOG_FILE}.txt /${LOG_FILE}.txt
COPY .env /.env
COPY . .

RUN go build -o /bareksa-interview

### Deploy built binary

FROM gcr.io/distroless/base-debian10

ARG LOG_FILE

WORKDIR /usr/src/app

COPY --from=build /${LOG_FILE}.txt ./${LOG_FILE}.txt
COPY --from=build /.env ./.env
COPY --from=build /bareksa-interview ./bareksa-interview

# USER nonroot:nonroot

ENTRYPOINT ["./bareksa-interview"]