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
	Result interface{} `json:"result"`
}

type responseError struct {
	response
	Errors []validator.Error `json:"errors,omitempty"`
}

func JSON(c *fiber.Ctx, data interface{}, status ...int) error {
	var s = http.StatusOK
	if len(status) > 0 {
		s = status[0]
	}

	_, ok := messages[s]
	if !ok {
		log.Printf(`[WARN] No message preset for status %v, proceeding with "Success" as response message and status code`, s)
		s = http.StatusOK
	}

	return c.Status(s).JSON(responseResult{
		response: response{
			Message: messages[s],
		},
		Result: data,
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
