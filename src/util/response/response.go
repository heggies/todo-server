package response

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/heggies/todo-server/src/util/validator"
)

type response struct {
	Message string `json:"message"`
}

type responseResult struct {
	response
	Result interface{} `json:"result,omitempty"`
}

type responseError struct {
	response
	Errors []validator.Error `json:"errors,omitempty"`
}

func JSON(c *fiber.Ctx, status int, data ...interface{}) error {
	_, ok := messages[status]
	if !ok {
		log.Printf(`[WARN] No message preset for status %v, proceeding with "Success" as response message and status code`, status)
		status = http.StatusOK
	}

	var result interface{} = data
	if len(data) == 0 {
		result = nil
	} else if len(data) == 1 {
		result = data[0]
	}

	return c.Status(status).JSON(responseResult{
		response: response{
			Message: messages[status],
		},
		Result: result,
	})
}

func Error(c *fiber.Ctx, status int, err ...validator.Error) error {
	_, ok := messages[status]
	if !ok {
		log.Printf(`[WARN] No message preset for status %v, proceeding with "Internal Server Error" as response message and status code`, status)
		status = http.StatusInternalServerError
	}

	return c.Status(status).JSON(responseError{
		response: response{
			Message: messages[status],
		},
		Errors: err,
	})
}
