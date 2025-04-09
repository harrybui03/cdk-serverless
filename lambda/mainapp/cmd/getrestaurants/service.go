package main

import (
	"cdk/lambda/mainapp/entities"
	"context"
)

type IRepository interface {
	GetRestaurants(ctx context.Context) ([]entities.Restaurant, error)
}

type usecase struct {
	repo IRepository
}

func (u usecase) GetRestaurants(ctx context.Context) ([]entities.Restaurant, error) {
	restaurant, err := u.repo.GetRestaurants(ctx)
	if err != nil {
		return []entities.Restaurant{}, err
	}

	return restaurant, nil
}
