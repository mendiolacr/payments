FROM golang:latest 

ADD ./ /src/app
WORKDIR /src/app
RUN go build -o payment_platform . 

CMD ["./payment_platform"]
