package main

import (
	"cdk/app"
	"cdk/lambda/mainapp/repository"
	"cdk/utils"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var db *dynamodb.DynamoDB

func init() {
	db = repository.GetDynamoDB()
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["restaurantId"]

	u := usecase{
		repo: &repository.Repository{
			Db: db,
		},
	}

	if err := u.DeleteRestaurant(ctx, id); err != nil {
		return utils.ErrorResponse(app.NewInternalError(err, err.Error())), nil
	}

	return utils.SimpleSuccessResponse("Create restaurant successfully"), nil
}

func main() {
	lambda.Start(handler)
}
