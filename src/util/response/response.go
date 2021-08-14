package response

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type response struct {
	Result  interface{} `json:"result"`
	Message string      `json:"message"`
}

var messages = map[int]string{
	http.StatusOK:                  "Success",
	http.StatusBadRequest:          "Bad Request",
	http.StatusInternalServerError: "Internal Server Error",
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

	return c.Status(s).JSON(response{
		Result:  data,
		Message: messages[s],
	})
}
