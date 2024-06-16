package presentation

import (
	"time"
)

//監修に合ったパッケージ名かどうかは謎

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
// Field名、JSON絡む時は大文字始まりにしとくのが無難そう　https://vtc.hatenablog.com/entry/2022/03/27/120505