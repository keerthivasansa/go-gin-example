FROM golang:buster

WORKDIR /app

COPY * ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /mux-tester

EXPOSE 5000

CMD [ "mux-tester" ]
