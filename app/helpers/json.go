package helpers

import (
	"encoding/json"
	"net/http"
)

func ToRequestBody(req *http.Request, result interface{}) {
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteResponseBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfError(err)
}
