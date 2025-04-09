package main

import (
	"cdk/lambda/mainapp/entities"
	"context"
)

type IRepository interface {
	GetRestaurantByID(ctx context.Context, id string) (entities.Restaurant, error)
	DeleteRestaurant(ctx context.Context, id string) error
}

type usecase struct {
	repo IRepository
}

func (u usecase) DeleteRestaurant(ctx context.Context, id string) error {
	restaurantToDelete, err := u.repo.GetRestaurantByID(ctx, id)
	if err != nil {
		return err
	}

	if err := u.repo.DeleteRestaurant(ctx, restaurantToDelete.Id); err != nil {
		return err
	}

	return nil
}
