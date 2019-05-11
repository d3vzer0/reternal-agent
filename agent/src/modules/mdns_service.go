package modules

import (
	"encoding/base64"
	"os"
	"os/signal"
	"syscall"

	"github.com/grandcat/zeroconf"
)

func MDNSService(input string) string {
	server, err := zeroconf.Register("microsoft-ds", "_microsoft-ds._tcp", "local.", 445, []string{"txtv=0", "lo=1", "la=2"}, nil)
	if err != nil {
		result := "Unable to run mDNS server"
		return base64.StdEncoding.EncodeToString([]byte(result))
	}
	defer server.Shutdown()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	select {
	case <-sig:
	}

	result := "Succesfully stopped service"
	return base64.StdEncoding.EncodeToString([]byte(result))
}
