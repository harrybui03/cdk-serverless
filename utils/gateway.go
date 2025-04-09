package utils

import (
	"cdk/app"
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

func Bind(request events.APIGatewayProxyRequest, data interface{}) error {
	err := json.Unmarshal([]byte(request.Body), data)
	if err != nil {
		return app.NewBadRequestError(err, "Invalid Request")
	}

	return nil
}

func SimpleSuccessResponse(message string) events.APIGatewayProxyResponse {
	payload := map[string]string{
		"message": message,
	}
	j, err := json.Marshal(payload)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       `{"message": "internal error"}`,
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(j),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		},
	}
}

func SuccessResponse(message string, data interface{}) events.APIGatewayProxyResponse {
	payload := map[string]interface{}{
		"message": message,
		"data":    data,
	}
	j, err := json.Marshal(payload)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       `{"message": "internal error"}`,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
		}
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(j),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		},
	}
}

func ErrorResponse(err error) events.APIGatewayProxyResponse {
	var val *app.Error
	ok := errors.As(err, &val)
	if ok {
		res := map[string]string{
			"message": val.Error(),
		}
		j, err := json.Marshal(res)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       `{"message": "internal error"}`,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
			}
		}

		return events.APIGatewayProxyResponse{
			StatusCode: val.Code,
			Body:       string(j),
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       `{"message": "internal error"}`,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		},
	}
}
