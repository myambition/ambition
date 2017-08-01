FROM golang:1.8.3

WORKDIR /go/src/github.com/myambition/ambition

ADD vendor vendor
ADD internal internal
ADD services services

ARG service=model

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./services/$service/$service-service/cmd/$service-server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/myambition/ambition/app .
CMD ["./app"]
