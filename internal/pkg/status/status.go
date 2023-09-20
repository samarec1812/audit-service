package status

import "net/http"

func GetStringStatusByCode(status int) string {
	switch status {
	case http.StatusAccepted, http.StatusOK:
		return "OK"
	case http.StatusBadRequest:
		return "Bad request"
	case http.StatusInternalServerError:
		return "Server error"
	default:
		return "Other error"
	}
}
