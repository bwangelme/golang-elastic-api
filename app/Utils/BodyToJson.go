package Utils

import (
	"encoding/json"
	"net/http"
)

func BodyToJson(r *http.Request) map[string]interface{} {
	if r.Body == nil {
		return nil
	}
	decoder := json.NewDecoder(r.Body)
	var dat map[string]interface{}
	err := decoder.Decode(&dat)
	if err != nil {
		panic(err)
	}
	return dat
}
