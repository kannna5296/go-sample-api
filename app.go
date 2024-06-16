package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
	"go-sample-api/presentation"
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
	res := presentation.BookDetailResponse{
		Id: "1", 
		Name:"hoge",
		Author: "HOGEHOGE", 
		CanRental: true, 
		Rentals: []presentation.BookDetailRentalResponse {
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