# go-practice

## 準備

.envファイルを作成し、DBとの接続情報を入力(SOTAkkkkに聞く)
```env
# .env
```

## 実行
```shell
go run .
```

## Linux用バイナリにコンパイル
```shell
GOOS=linux GOARCH=amd64 go build -o go-practice main.go
```
```shell
chmod +x go-practice
```

## GCPデプロイ
初期化
```shell
gcloud init  
```

App engineにデプロイ
```shell
gcloud app deploy
```
