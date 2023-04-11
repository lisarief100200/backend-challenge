package shop

import (
	"time"
)

type Customer struct {
	ID              int
	CustomerName    string
	CustomerContNo  string
	CustomerAddress string
	TotalBuy        int
	CreatorId       int
	CreatedAt       time.Time
}
