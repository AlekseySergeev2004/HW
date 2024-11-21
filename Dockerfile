FROM golang:1.20

RUN mkdir -p /usr/src/app/
WORKDIR /usr/src/app/

COPY go.mod /usr/src/app/

RUN go mod download

COPY . /usr/src/app/

RUN go build -o main .

EXPOSE 32777

CMD ["./main"]


