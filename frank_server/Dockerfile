FROM golang:1.13.15-alpine3.12
RUN echo test
WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /run_server ./cmd/api

EXPOSE 8088

CMD [ "/run_server" ]
