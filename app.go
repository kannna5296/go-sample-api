package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

func main() {
  // インスタンスを作成
  e := echo.New()

  // ミドルウェアを設定
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // ルートを設定
  e.GET("/book",bookDetail)

  // サーバーをポート番号1323で起動
  e.Logger.Fatal(e.Start(":1323"))
}

func bookDetail(c echo.Context) error {
	// DBとかイイので一旦ゴリ押してResponseをJSON表現
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
	//ここでログを出すには？？
    return c.JSON(http.StatusOK, res)
}

// Field名、JSON絡む時は大文字始まりにしとくのが無難そう　https://vtc.hatenablog.com/entry/2022/03/27/120505

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