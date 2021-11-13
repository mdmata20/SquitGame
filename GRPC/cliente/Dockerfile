FROM golang
WORKDIR /
COPY . .
#ENV HOST=0.0.0.0:8082
RUN go mod download
EXPOSE 8080
CMD ["go", "run", "cliente.go"]