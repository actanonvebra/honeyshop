FROM golang:1.22-alpine AS builder

#Calısacagı sanal yer.
WORKDIR /app

COPY .env /app/.env

#localdeki tüm dosyaları olusturdugum sanal /app icerisine atar.
COPY . .

#go'nun bağımlılıklarını halletmek icin var.
RUN go mod tidy

#derledik.
RUN go build -o main ./cmd 


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]

