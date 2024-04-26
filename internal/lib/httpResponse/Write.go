package httpResponse

import "net/http"

func Write(w http.ResponseWriter, statusCode int, info string) {
	_, err := w.Write([]byte(info))
	if err != nil {
		return
	}
	w.WriteHeader(statusCode)
}
