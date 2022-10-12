# Builder
FROM golang:1.14.2-alpine3.11 as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /cmd/api

COPY . .

RUN make engine

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /cmd/api 

WORKDIR /cmd/api 

EXPOSE 9030

COPY --from=builder /cmd/api/engine /cmd/api

CMD /cmd/api/engine