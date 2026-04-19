# ベースイメージとして公式の Golang イメージを使用
FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./
COPY src ./src

RUN go build -o myapp ./src

CMD ["./myapp"]
