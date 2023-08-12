FROM golang:1.20-alpine

ENV PORT=8080

ENV HOST=0.0.0.0

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN  go build -o app *.go

EXPOSE 8080

CMD [ "./app" ]