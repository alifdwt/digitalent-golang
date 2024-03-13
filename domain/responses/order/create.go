package order

import "time"

type ItemResponse struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
}

type OrderResponse struct {
	ID           uint           `json:"id"`
	OrderedAt    time.Time      `json:"orderedAt"`
	CustomerName string         `json:"customerName"`
	Items        []ItemResponse `json:"items"`
}
