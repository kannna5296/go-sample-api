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

## Docker まわり

### mysql ログイン

```
docker exec -ti docker_mysql mysql -u root -p
```

### initdb 流し直したい時

```
docker-compose down --volume //Volumeも消す

```




・スライスと配列の違い
・VSCodeでの開発体験について（Golandのようが良い？
・Field名、大文字始まりにしとくのが無難そう　https://vtc.hatenablog.com/entry/2022/03/27/120505

・「クラス」と「構造体」の違い


## JSON構造体の定義

```
	res := BookDetailResponse{
		Id: "1", 
		Name:"hoge",
		Author: "HOGEHOGE", 
		CanRental: true, 
		Rentals: []BookDetailRentalResponse {
			{
				UserId: "11",
				RentedAt: time.Now(), 
				Deadline: time.Now(), 
				Returned: false,
			},
		},
	}

  type BookDetailResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
	CanRental bool `json:"canRental"`
	Rentals []BookDetailRentalResponse `json:"rentals"`
}

type BookDetailRentalResponse struct {
	UserId string `json:"userId"`
	RentedAt time.Time `json:"rentedAt"`
	Deadline time.Time `json:"deadline"`
	Returned bool `json:"returned"`
}
```