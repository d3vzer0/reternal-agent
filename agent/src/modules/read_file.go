package modules

import (
	"encoding/base64"
	"io/ioutil"
)

func ReadFile(path string) string {
	file_content, err := ioutil.ReadFile(path)
	result := ""
	if err != nil {
		result = err.Error()
	} else {
		result = base64.StdEncoding.EncodeToString(file_content)
	}
	return result
}
