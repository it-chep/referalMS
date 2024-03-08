FROM golang:alpine
RUN apk add ffmpeg
WORKDIR /app
COPY . .
RUN go mod download

EXPOSE 8000
CMD ["go", "run", "cmd/referal-server/local/main.go"]
