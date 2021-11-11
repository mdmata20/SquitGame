FROM golang

WORKDIR /

COPY . .

RUN go mod download

EXPOSE 80

CMD ["go","run","Sender.go"]