// +build darwin

package mapping

import (
	Modules "../modules"
)

func FunctionMapping() map[string]func(string) string {
	var function_mapping = map[string]func(string) string{
		"exec_shell":      Modules.ExecShell,
		"get_ifaces":      Modules.GetIfaces,
		"download_remote": Modules.DownloadRemote,
		"read_file":       Modules.ReadFile,
		"mdns_service":    Modules.MDNSService,
	}

	return function_mapping

}
