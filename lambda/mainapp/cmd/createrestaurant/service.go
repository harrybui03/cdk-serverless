package main

import (
	"cdk/lambda/mainapp/entities"
	"context"
	"time"
)

type IRepository interface {
	UpsertRestaurant(ctx context.Context, restaurant entities.Restaurant) error
}

type usecase struct {
	repo IRepository
}

func (u usecase) CreateRestaurant(ctx context.Context, input CreateRestaurantDTO) error {
	now := time.Now().Unix()
	restaurantEntity := entities.Restaurant{
		Name:                 input.Name,
		Address:              input.Address,
		CuisineType:          input.CuisineType,
		Rating:               input.Rating,
		PhoneNumber:          input.PhoneNumber,
		OpeningHours:         input.OpeningHours,
		IsVegetarianFriendly: input.IsVegetarianFriendly,
		CreatedAt:            now,
		UpdatedAt:            now,
	}

	if err := u.repo.UpsertRestaurant(ctx, restaurantEntity); err != nil {
		return err
	}

	return nil
}
