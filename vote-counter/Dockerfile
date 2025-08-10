FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go mod init vote-counter && go get
RUN go build -o vote-counter
CMD ["./vc"]