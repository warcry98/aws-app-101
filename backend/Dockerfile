FROM golang:1.23 as build

ADD . /go/src/
WORKDIR /go/src/
RUN go get
RUN go install

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/api

FROM alpine as prod
RUN apk update && apk add curl
WORKDIR /go/src/
COPY --from=build /go/bin/api /go/bin/api
COPY --from=build /go/src/.env /go/src/.env
EXPOSE 8080

ENTRYPOINT ["/go/bin/api"]
