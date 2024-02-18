# ベースイメージを指定
FROM golang:latest

# アプリケーションのソースコードを追加
ADD . /app

# ワーキングディレクトリを設定
WORKDIR /app

# アプリケーションのビルド
RUN go build -o main .

# ポート8081を公開
EXPOSE 8081

# アプリケーションの実行
CMD ["./main"]
