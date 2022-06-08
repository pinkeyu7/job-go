package model

import "time"

type BillingUsageHour struct {
	Id         int       `json:"id"`
	Key        string    `json:"key"`
	TimeByHour time.Time `json:"time_by_hour"`
	ServiceId  int       `json:"service_id"`
	SkuId      int       `json:"sku_id"`
	PriceId    int       `json:"price_id"`
	Price      string    `json:"price"`
	Usage      string    `json:"usage"`
	Cost       string    `json:"cost"`
	CreatedAt  time.Time `json:"created_at" xorm:"created"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"updated"`
}
