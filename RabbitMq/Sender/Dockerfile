FROM golang

WORKDIR /

COPY . .

RUN go mod download

EXPOSE 3005

CMD ["go","run","Sender.go"]

