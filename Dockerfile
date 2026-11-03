FROM alpine:latest

COPY / /audiogo

WORKDIR /audiogo

RUN go mod download

RUN go build -o bin/audiogo

CMD [ "bin/audiogo" ]