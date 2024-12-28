# ベースイメージ
FROM golang:1.23

# 作業ディレクトリを設定
WORKDIR /app

# 必要なファイルをコピー
COPY go.mod ./
RUN go mod download
COPY . .

# アプリケーションをビルド
RUN go build -o /docker_test main.go

# ポートを公開
EXPOSE 8000

# アプリケーションを実行
CMD ["/docker_test"]
