#@version 3.7.5
FROM golang:1.14.6-alpine3.12 as base

WORKDIR /go/src/
COPY ./go/src ./go_hpa

WORKDIR /go/src/go_hpa

RUN go get -v ./...  \
    && go build -ldflags="-s -w"
RUN ls -la

CMD ["go_hpa"]

#FROM scratch as prod
FROM alpine as prod

COPY --from=base /go/src/go_hpa/go_hpa .

CMD ["/go_hpa"]
