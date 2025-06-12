package entities

type Product struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	KG       *int64 `json:"kg,omitempty"`
	Quantity *int64 `json:"quantity,omitempty"`
}
