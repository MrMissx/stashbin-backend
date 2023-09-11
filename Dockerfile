FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go mod verify

COPY . .

RUN go build -o stashbin


FROM golang:1.21-alpine

WORKDIR /app

ENV GIN_MODE=release

COPY --from=build /app/stashbin .

CMD ["./stashbin"]
