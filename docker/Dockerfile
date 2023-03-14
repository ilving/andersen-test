FROM golang:1.19-alpine as builder

RUN apk update && apk add --no-cache git openssh-client

COPY go.mod /app
COPY go.sum /app
RUN go mod download && go mod verify
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/app/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
