FROM golang

WORKDIR /

COPY . .

RUN go mod download

EXPOSE 2062

CMD ["go","run","Sender.go"]

