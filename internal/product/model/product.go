package model

type Product struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
	Quantity    int32    `json:"quantity"`
	Price       float64  `json:"price"`
	Version     int32    `json:"version"`
}
