FROM golang:1.10-alpine as build-env
RUN apk add --update curl build-base git
RUN curl https://glide.sh/get | sh
WORKDIR /go/src/github.com/joshrendek/the-counter
COPY glide.yaml /go/src/github.com/joshrendek/the-counter
COPY glide.lock /go/src/github.com/joshrendek/the-counter
RUN glide install
ADD . /go/src/github.com/joshrendek/the-counter
RUN make build

FROM alpine:3.7
RUN apk add --update curl jq
WORKDIR /app
ENV PORT 8080
COPY --from=build-env /go/src/github.com/joshrendek/the-counter/api /app/api
COPY --from=build-env /go/src/github.com/joshrendek/the-counter/helm-test.sh /app/helm-test.sh
ENTRYPOINT ["./api"]
