package entities

type Restaurant struct {
	Id                   string            `json:"id"`
	Name                 string            `json:"name"`
	Address              string            `json:"address"`
	CuisineType          string            `json:"cuisine_type"`
	Rating               float64           `json:"rating"`
	PhoneNumber          string            `json:"phone_number"`
	OpeningHours         map[string]string `json:"opening_hours"`
	IsVegetarianFriendly bool              `json:"is_vegetarian_friendly"`
	CreatedAt            int64             `json:"created_at"`
	UpdatedAt            int64             `json:"updated_at"`
}
