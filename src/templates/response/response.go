package response

import (
	errors "dana/anabul-rest-api/src/templates/error"
	"encoding/json"
	"net/http"
)

type Response struct {
	Message *string      `json:"message,omitempty"`
	Data    *interface{} `json:"data,omitempty"`
}

func Json(response http.ResponseWriter, httpCode int, message string, data interface{}) {
	response.Header().Set("Content-Type", "application/json")
	// response.Header().Set("Access-Control-Allow-Origin", "*")
	// response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	// response.Header().Set("Access-Control-Allow-Methods", "*")

	response.WriteHeader(httpCode)
	res := Response{
		Message: &message,
		Data:    &data,
	}
	json.NewEncoder(response).Encode(res)
}

func Success(response http.ResponseWriter, data interface{}) {
	response.Header().Set("Content-Type", "application/json")
	// response.Header().Set("Access-Control-Allow-Origin", "*")
	// response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	// response.Header().Set("Access-Control-Allow-Methods", "*")

	response.WriteHeader(http.StatusOK)
	message := "SUCCESS"
	res := Response{
		Message: &message,
		Data:    &data,
	}
	json.NewEncoder(response).Encode(res)
}

func Error(response http.ResponseWriter, err error) {
	_, ok := err.(*errors.RespError)
	if !ok {
		err = errors.InternalServerError(err.Error())
	}

	er, _ := err.(*errors.RespError)
	response.Header().Set("Content-Type", "application/json")
	// response.Header().Set("Access-Control-Allow-Origin", "*")
	// response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	// response.Header().Set("Access-Control-Allow-Methods", "*")

	response.WriteHeader(er.Code)
	res := Response{
		Message: &er.Message,
	}
	json.NewEncoder(response).Encode(res)
}
