package response

import "net/http"

var messages = map[int]string{
	http.StatusOK:                  "Success",
	http.StatusBadRequest:          "Bad Request",
	http.StatusNotFound:            "Not Found",
	http.StatusInternalServerError: "Internal Server Error",
}
