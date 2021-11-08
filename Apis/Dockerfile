FROM golang

WORKDIR /

COPY . .

RUN go mod download

EXPOSE 3010

CMD ["go","run","main.go"]