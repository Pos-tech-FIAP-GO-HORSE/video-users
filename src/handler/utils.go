package handler

import (
	"encoding/json"

	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/core/responses"
	"github.com/aws/aws-lambda-go/events"
)

var headers = map[string]string{
	"Content-Type": "application/json",
}

func apiGatewayResponse(statusCode int, body string) events.APIGatewayV2HTTPResponse {
	return events.APIGatewayV2HTTPResponse{
		Headers:    headers,
		StatusCode: statusCode,
		Body:       body,
	}
}

func apiGatewayResponseWithError(statusCode int, body string) events.APIGatewayV2HTTPResponse {
	errorResponse := responses.Error{
		Error: body,
	}

	errorResponseRaw, _ := json.Marshal(errorResponse)

	return events.APIGatewayV2HTTPResponse{
		Headers:    headers,
		StatusCode: statusCode,
		Body:       string(errorResponseRaw),
	}
}
