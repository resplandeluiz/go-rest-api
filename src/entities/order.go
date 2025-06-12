package entities

type Order struct {
	ID       int64     `json:"id"`
	ONGID    string    `json:"ong_id"`
	Products []product `json:"produtos"`
}
