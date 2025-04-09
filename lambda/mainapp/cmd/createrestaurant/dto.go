package main

type CreateRestaurantDTO struct {
	Name                 string            `json:"name" binding:"required"`
	Address              string            `json:"address" binding:"required"`
	CuisineType          string            `json:"cuisine_type" binding:"required"`
	Rating               float64           `json:"rating"`
	PhoneNumber          string            `json:"phone_number"`
	OpeningHours         map[string]string `json:"opening_hours"`
	IsVegetarianFriendly bool              `json:"is_vegetarian_friendly"`
}
