package core

import (
	"log"
)

const (
	STATUS_SUCCESS            = 1
	STATUS_INVALID_LOGIN      = 3
	STATUS_INVALID_PARAMETERS = 4
)

var STATUS_MESSAGE = map[uint]string{
	1: "Success",
	4: "Invalid Parameters",
}

type ApiResponse struct {
	status  uint
	message string
	data    map[string]interface{}
	token   map[string]interface{}
}

/**
 * @brief      Sets the status.
 * @param      status  The status
 * @return     object
 */
func (api *ApiResponse) SetStatus(status uint) {
	api.status = status
	api.message = STATUS_MESSAGE[status]
}

/**
 * @brief      Sets the message.
 * @param      message  The message
 * @return     object
 */
func (api *ApiResponse) SetMessage(message string) {
	api.message = message
}

/**
 * @brief      Sets the data.
 * @param      data  The data
 * @return     object
 */
func (api *ApiResponse) SetData(data map[string]interface{}) {
	api.data = data
}

/**
 * @brief      Sets the token.
 * @param      tokendata  The tokendata
 * @return     object
 */
func (api *ApiResponse) SetToken(tokendata map[string]interface{}) {
	api.token = tokendata
}

/**
 * @brief      this function for to send response
 * @return     object
 */
func (api *ApiResponse) ToResponse() map[string]interface{} {
	responseData := map[string]interface{}{
		"data":    api.data,
		"status":  api.status,
		"message": api.message,
		"token":   api.token,
	}
	log.Println(responseData)

	api.data = nil

	return responseData
}
