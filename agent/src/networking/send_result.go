package networking

import (
	"bytes"
	"net/http"
)

func SendResult(base_url string, task_result []byte) {
	// Send core agent data to back-end API
	response, response_error := http.Post(base_url, "application/json", bytes.NewBuffer(task_result))
	if response_error != nil {
		_ = response_error
		_ = response
	}
}
