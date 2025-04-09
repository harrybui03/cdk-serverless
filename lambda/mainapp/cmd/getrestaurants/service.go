package main

import (
	"cdk/lambda/mainapp/entities"
	"context"
)

type IRepository interface {
	GetAllRestaurants(ctx context.Context) ([]entities.Restaurant, error)
}

type usecase struct {
	repo IRepository
}

func (u usecase) GetRestaurants(ctx context.Context) ([]entities.Restaurant, error) {
	restaurant, err := u.repo.GetAllRestaurants(ctx)
	if err != nil {
		return []entities.Restaurant{}, err
	}

	return restaurant, nil
}
