package networking

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func SendPulse(beacon_data []byte, base_url string) interface{} {
	// Send core agent data to back-end API
	response, response_error := http.Post(base_url, "application/json", bytes.NewBuffer(beacon_data))
	if response_error != nil {
		_ = response_error
	}
	// Parse HTTP response content
	response_content, read_error := ioutil.ReadAll(response.Body)
	if read_error != nil {
		_ = read_error
	}
	// Parse response content as JSON and convert to interface mapping
	var json_values interface{}
	json_response := json.Unmarshal(response_content, &json_values)
	if json_response != nil {
		_ = json_response
	}
	return json_values
}
