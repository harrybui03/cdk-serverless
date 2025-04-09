package register

import (
	"cdk/utils"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/jsii-runtime-go"
)

type DynamoDBTable struct {
	TableRestaurant awsdynamodb.Table
}

func MainApp(stack awscdk.Stack) {
	createRestaurantHandler := awscdklambdagoalpha.NewGoFunction(stack, utils.JsiiWithEnv("main_app_create_restaurant_handler"), &awscdklambdagoalpha.GoFunctionProps{
		Entry:        jsii.String("./lambda/mainapp/cmd/createrestaurant"),
		FunctionName: jsii.String(utils.NameWithEnv("main_app_create_restaurant_handler")),
	})
	deleteRestaurantHandler := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("main_app_delete_restaurant_handler"), &awscdklambdagoalpha.GoFunctionProps{
		Entry:        jsii.String("./lambda/mainapp/cmd/deleterestaurant"),
		FunctionName: jsii.String("main_app_delete_restaurant_handler"),
	})

	getRestaurantByIdHandler := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("main_app_get_restaurant_by_id_handler"), &awscdklambdagoalpha.GoFunctionProps{
		Entry:        jsii.String("./lambda/mainapp/cmd/getrestaurantbyid"),
		FunctionName: jsii.String("main_app_get_restaurant_by_id_handler"),
	})

	getRestaurantsHandler := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("main_app_get_restaurants_handler"), &awscdklambdagoalpha.GoFunctionProps{
		Entry:        jsii.String("./lambda/mainapp/cmd/getrestaurants"),
		FunctionName: jsii.String("main_app_get_restaurants_handler"),
	})

	updateRestaurantHandler := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("main_app_update_restaurant_handler"), &awscdklambdagoalpha.GoFunctionProps{
		Entry:        jsii.String("./lambda/mainapp/cmd/updaterestaurant"),
		FunctionName: jsii.String("main_app_update_restaurant_handler"),
	})

	table := RegisterTable(stack)
	table.TableRestaurant.GrantFullAccess(createRestaurantHandler)
	table.TableRestaurant.GrantFullAccess(deleteRestaurantHandler)
	table.TableRestaurant.GrantFullAccess(getRestaurantByIdHandler)
	table.TableRestaurant.GrantFullAccess(getRestaurantsHandler)
	table.TableRestaurant.GrantFullAccess(updateRestaurantHandler)

	httpApi := awsapigatewayv2.NewHttpApi(stack, utils.JsiiWithEnv("main_app_http_api"), &awsapigatewayv2.HttpApiProps{
		ApiName: jsii.String(utils.NameWithEnv("main_app_http_api")),
		CorsPreflight: &awsapigatewayv2.CorsPreflightOptions{
			AllowHeaders: &[]*string{
				jsii.String("*"),
			},
			AllowMethods: &[]awsapigatewayv2.CorsHttpMethod{
				awsapigatewayv2.CorsHttpMethod_ANY,
				awsapigatewayv2.CorsHttpMethod_DELETE,
				awsapigatewayv2.CorsHttpMethod_GET,
				awsapigatewayv2.CorsHttpMethod_HEAD,
				awsapigatewayv2.CorsHttpMethod_OPTIONS,
				awsapigatewayv2.CorsHttpMethod_PATCH,
				awsapigatewayv2.CorsHttpMethod_PUT,
				awsapigatewayv2.CorsHttpMethod_POST,
			},
			AllowOrigins: &[]*string{
				jsii.String("*"),
			},
		},
	})

	httpApi.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path: jsii.String("/restaurants"),
		Methods: &[]awsapigatewayv2.HttpMethod{
			awsapigatewayv2.HttpMethod_POST,
		},
		Integration: awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("mainAppCreateRestaurantIntegration"), createRestaurantHandler, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{}),
	})

	httpApi.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path: jsii.String("/restaurants/{id}"),
		Methods: &[]awsapigatewayv2.HttpMethod{
			awsapigatewayv2.HttpMethod_DELETE,
		},
		Integration: awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("mainAppDeleteRestaurantIntegration"), deleteRestaurantHandler, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{}),
	})

	httpApi.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path: jsii.String("/restaurants/{id}"),
		Methods: &[]awsapigatewayv2.HttpMethod{
			awsapigatewayv2.HttpMethod_GET,
		},
		Integration: awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("mainAppGetRestaurantByIdIntegration"), getRestaurantByIdHandler, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{}),
	})

	httpApi.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path: jsii.String("/restaurants"),
		Methods: &[]awsapigatewayv2.HttpMethod{
			awsapigatewayv2.HttpMethod_GET,
		},
		Integration: awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("mainAppGetRestaurantsIntegration"), getRestaurantsHandler, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{}),
	})

	httpApi.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path: jsii.String("/restaurants/{id}"),
		Methods: &[]awsapigatewayv2.HttpMethod{
			awsapigatewayv2.HttpMethod_PUT,
		},
		Integration: awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("mainAppUpdateRestaurantIntegration"), updateRestaurantHandler, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{}),
	})

}

func RegisterTable(stack awscdk.Stack) DynamoDBTable {
	tableRestaurant := awsdynamodb.NewTable(stack, utils.JsiiWithEnv("table_restaurant"), &awsdynamodb.TableProps{
		TableName: jsii.String(utils.NameWithEnv("main_app_restaurant")),
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		BillingMode: awsdynamodb.BillingMode_PAY_PER_REQUEST,
	})

	return DynamoDBTable{
		TableRestaurant: tableRestaurant,
	}
}
