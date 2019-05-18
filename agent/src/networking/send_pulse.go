package networking

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func SendPulse(beacon_data []byte, base_url string, public_key_string string) interface{} {
	// Send core agent data to back-end API
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	encoded_message, aes_key := EncodeMessage(beacon_data, public_key_string)
	response, response_error := client.Post(base_url, "application/json", bytes.NewBuffer(encoded_message))
	if response_error != nil {
		return response_error
	}
	response_content, read_error := ioutil.ReadAll(response.Body)
	if read_error != nil {
		return read_error
	}
	decoded_response, err := base64.StdEncoding.DecodeString(string(response_content))
	if err != nil {
		return err
	}
	decrypted_response := DecryptMessage(aes_key, decoded_response)
	var json_values interface{}
	json_response := json.Unmarshal(decrypted_response, &json_values)
	if json_response != nil {
		return json_response
	}
	return json_values

}
