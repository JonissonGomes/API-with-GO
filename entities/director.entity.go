package entities

type Director struct {
	Firstname          string `json:"firstname"`
	Lastname           string `json:"lastname"`
	QuantityProduction int    `json:"quantityProduction"`
}
