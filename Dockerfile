FROM golang:1.14

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o realstate.so ./main.go

EXPOSE 8080

ENTRYPOINT [ "./realstate.so" ]