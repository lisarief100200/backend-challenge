package api

type CustomerResponse struct {
	ID              int    `json:"id"`
	CustomerName    string `json:"customer_name"`
	CustomerContNo  string `json:"customer_cont_no"`
	CustomerAddress string `json:"customer_address"`
	TotalBuy        int    `json:"total_buy"`
}
