package modules

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func DownloadRemote(input string) string {
	arguments := strings.Fields(input)
	response, err := http.Get(arguments[0])
	result := ""
	if err != nil {
		result = err.Error()
	} else {
		defer response.Body.Close()
		out, err := os.Create(arguments[1])
		if err != nil {
			result = err.Error()
		} else {
			_, err := io.Copy(out, response.Body)
			_ = err
			string_output := fmt.Sprintf("Downloaded %s to %s", arguments[0], arguments[1])
			result = base64.StdEncoding.EncodeToString([]byte(string_output))
		}
	}
	return result
}
