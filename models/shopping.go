package api

import "encoding/json"

type CustomerRequest struct {
	CustomerName    string      `json:"customer_name" binding:"required"`
	CustomerContNo  string      `json:"customer_cont_no" binding:"required"`
	CustomerAddress string      `json:"customer_address" binding:"required"`
	TotalBuy        json.Number `json:"total_buy" binding:"required"`
	CreatorId       json.Number `json:"creator_id" binding:"required"`
}
