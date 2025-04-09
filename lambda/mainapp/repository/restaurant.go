package repository

import (
	"cdk/lambda/mainapp/entities"
	"cdk/utils"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rs/zerolog/log"
)

func (r Repository) DeleteRestaurant(ctx context.Context, id string) error {
	log.Info().Str("table", utils.NameWithEnv("restaurants")).Msg("Delete Restaurant")
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(utils.NameWithEnv("restaurants")),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	_, err := r.Db.DeleteItemWithContext(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) UpsertRestaurant(ctx context.Context, restaurant entities.Restaurant) error {
	log.Info().Str("table", utils.NameWithEnv("restaurants")).Msg("Upsert Restaurant")
	av, err := dynamodbattribute.MarshalMap(restaurant)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(utils.NameWithEnv("restaurants")),
		Item:      av,
	}

	_, err = r.Db.PutItemWithContext(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) GetAllRestaurants(ctx context.Context) ([]entities.Restaurant, error) {
	log.Info().Str("table", utils.NameWithEnv("restaurants")).Msg("Scanning Restaurants")
	input := &dynamodb.ScanInput{
		TableName: aws.String(utils.NameWithEnv("restaurants")),
	}

	result, err := r.Db.ScanWithContext(ctx, input)
	if err != nil {
		log.Error().Err(err).Msg("Error scanning restaurants table")
		return []entities.Restaurant{}, err
	}

	var restaurants []entities.Restaurant
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &restaurants)
	if err != nil {
		log.Error().Err(err).Msg("Error unmarshalling scanned restaurant items")
		return []entities.Restaurant{}, err
	}

	log.Info().Int("count", len(restaurants)).Msg("Successfully scanned restaurants")
	return restaurants, nil
}

func (r Repository) GetRestaurantByID(ctx context.Context, id string) (entities.Restaurant, error) {
	log.Info().Str("table", utils.NameWithEnv("restaurants")).Msg("Get Restaurant By Id")
	input := &dynamodb.QueryInput{
		TableName:              aws.String(utils.NameWithEnv("restaurants")),
		KeyConditionExpression: aws.String("id = :id"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {
				S: aws.String(id),
			},
		},
	}

	result, err := r.Db.QueryWithContext(ctx, input)
	if err != nil {
		return entities.Restaurant{}, err
	}

	var restaurant entities.Restaurant
	err = dynamodbattribute.UnmarshalMap(result.Items[0], &restaurant)
	if err != nil {
		return entities.Restaurant{}, err
	}

	return restaurant, nil
}
