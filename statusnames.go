package httpjson

import "net/http"

var statusMessages map[int]string = map[int]string{
	http.StatusOK:                  "OK",
	http.StatusAccepted:            "ACCEPTED",
	http.StatusBadRequest:          "BAD_REQUEST",
	http.StatusUnauthorized:        "UNAUTHORIZED",
	http.StatusForbidden:           "FORBIDDEN",
	http.StatusNotFound:            "NOT_FOUND",
	http.StatusMethodNotAllowed:    "METHOD_NOT_ALLOWED",
	http.StatusInternalServerError: "INTERNAL_SERVER_ERROR",
	http.StatusTooManyRequests:     "TOO_MANY_REQUESTS",
}

func statusMessage(status int) (message string) {
	if val, ok := statusMessages[status]; ok {
		message = val
	} else {
		message = "UNKNOWN"
	}
	return
}
