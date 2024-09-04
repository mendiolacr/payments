FROM golang:latest 

ADD ./ /src/app
WORKDIR /src/app
RUN go build -o main . 

CMD ["./main"]