package util

import (
	"encoding/json"
	"reflect"
	"strconv"
)

type response struct {
	Result string      `json:"result"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data"`
}

func ReturnResponse(respData interface{}, status int, result, message string) (int, string, []byte) {
	rt := reflect.TypeOf(respData)
	var resp response

	switch rt.Kind() {
	case reflect.Slice, reflect.Array:
		respData, _ := json.Marshal(respData)
		resp = response{
			Result: result,
			Info:   message,
			Data:   json.RawMessage(respData),
		}
	case reflect.Int:

		resp = response{
			Result: result,
			Info:   message,
			Data:   json.RawMessage(strconv.Itoa(respData.(int))),
		}
	default:
		if value, ok := respData.(string); ok {
			respData, _ = json.Marshal(value)
		}
		resp = response{
			Result: result,
			Info:   message,
			Data:   json.RawMessage(respData.([]byte)),
		}
	}
	jsonResp, _ := json.MarshalIndent(resp, "", "    ")
	return status, "application/json", jsonResp
}
