package main

import (
	"cdk/lambda/mainapp/entities"
	"context"
)

type IRepository interface {
	GetRestaurantByID(ctx context.Context, id string) (entities.Restaurant, error)
}

type usecase struct {
	repo IRepository
}

func (u usecase) GetRestaurantByID(ctx context.Context, id string) (entities.Restaurant, error) {
	restaurant, err := u.repo.GetRestaurantByID(ctx, id)
	if err != nil {
		return entities.Restaurant{}, err
	}

	return restaurant, nil
}
