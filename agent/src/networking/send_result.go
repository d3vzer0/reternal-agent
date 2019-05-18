package networking

import (
	"bytes"
	"crypto/tls"
	"net/http"
)

func SendResult(base_url string, task_result []byte, public_key_string string) {
	// Send core agent data to back-end API
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	encoded_message, _ := EncodeMessage(task_result, public_key_string)
	response, response_error := client.Post(base_url, "application/json", bytes.NewBuffer(encoded_message))
	if response_error != nil {
		_ = response_error
		_ = response
	}
}
