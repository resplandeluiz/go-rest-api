package usecaseproduct

import "clean-arch-api/src/entities"

func RegisterProduct() entities.Product {
	return entities.Product{ID: 1, Name: "Arroz"}
}
