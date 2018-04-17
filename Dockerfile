FROM golang:1.10-alpine as build-env
RUN apk add --update curl build-base git
RUN curl https://glide.sh/get | sh
ADD . /go/src/github.com/joshrendek/the-counter
WORKDIR /go/src/github.com/joshrendek/the-counter
RUN make build

FROM alpine:3.7
WORKDIR /app
ENV PORT 8080
COPY --from=build-env /go/src/github.com/joshrendek/the-counter/api /app/api
ENTRYPOINT ["./api"]
