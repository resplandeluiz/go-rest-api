package entities

type product struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	KG       *int64 `json:"kg,omitempty"`
	Quantity *int64 `json:"quantity,omitempty"`
}
