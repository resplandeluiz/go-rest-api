package entities

type order struct {
	ID       int64     `json:"id"`
	ONGID    string    `json:"ong_id"`
	Products []product `json:"produtos"`
}
