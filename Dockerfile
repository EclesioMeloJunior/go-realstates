FROM golang:1.14

WORKDIR /app

COPY . .

RUN go build -o realstate ./main.go

EXPOSE 8080

ENTRYPOINT [ "./realstate --port ':8080'" ]