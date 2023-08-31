# go-sample-api

モジュールファイル初期化（ここにもろもろ依存関係が描かれる,build.gradle みたいなもん）

```
go mod init モジュール名
```

## Ver 管理

https://qiita.com/massaaaaan/items/1b4b494b78cc77c69ec8#goenv%E3%81%AE%E3%82%A4%E3%83%B3%E3%82%B9%E3%83%88%E3%83%BC%E3%83%AB

ここ参考にした

goenv 2.1.5 を利用

## ライブラリの理解

- net/http
  - https://pkg.go.dev/net/http
  - http クライアント（リクエストを投げる時用）
  - サーバ立て（リクエストを待ち受ける時用）
    - ルーティングには go-chi とかをくっつけて使う？
