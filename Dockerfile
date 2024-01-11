FROM golang:1.21.6-alpine as builder

WORKDIR /app

COPY go.mod .

ENV CGO_ENABLED 0

RUN go mod download

COPY . . 

RUN go build . 

FROM scratch

WORKDIR /

COPY --from=builder /app/filler /filler

ENTRYPOINT [ "/filler" ]

