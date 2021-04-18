FROM golang:1.16 AS build-env
WORKDIR /src

ADD ./go.mod .
RUN go mod download

ADD . .
RUN go build -o w0q

WORKDIR /app
RUN cp /src/w0q .


FROM debian
WORKDIR /app

ENV GIN_MODE=release
ENV W0Q_DB_PATH=/data/w0q.db

COPY --from=build-env /app .
COPY public/ /app/public

VOLUME /data

ENTRYPOINT ["/app/w0q"]