FROM golang:bullseye

# ワークディレクトリを設定
WORKDIR /app

# 依存関係を解決するために必要なファイルをコピー
COPY . /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
RUN go mod download \
    && go mod verify \
    && go build -v -o /app .

# アプリケーションをビルド
RUN go build main.go

# ポート8080を公開
EXPOSE 8080

# サーバー実行
CMD ["./main"]
