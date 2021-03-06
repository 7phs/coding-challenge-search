FROM golang:1.11-stretch

ADD . /go/src/github.com/7phs/coding-challenge-search
WORKDIR /go/src/github.com/7phs/coding-challenge-search

RUN make build

FROM debian:stretch

RUN apt-get update \
    && apt-get clean

EXPOSE 8080
WORKDIR /root/
COPY --from=0 /go/src/github.com/7phs/coding-challenge-search .
COPY fatlama.sqlite3 ./fatlama.sqlite3

ENV LOG_LEVEL info
ENV STAGE production
ENV ADDR :8080
ENV CORS true
ENV DB_URL ./fatlama.sqlite3

CMD ["./search-service", "run"]