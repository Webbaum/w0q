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
ENV PORT=8080
ENV W0Q_DB_PATH=/data/w0q.db
ENV W0Q_LEGAL_ADDRESS="add your address here"
ENV W0Q_LEGAL_MAIL="add your email address here"
ENV W0Q_LEGAL_PHONE="add your phone number here"

COPY --from=build-env /app .
COPY public/ /app/public

VOLUME /data

EXPOSE 8080

ENTRYPOINT ["/app/w0q"]
