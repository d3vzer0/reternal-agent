package modules

import (
	"crypto/sha1"
	"encoding/hex"
	"os"

	"github.com/denisbrodbeck/machineid"
)

func GenerateID(beacon_username string) string {
	// Generate unique beacon ID based on machine ID and Username
	machine_id, machine_error := machineid.ID()
	if machine_error != nil {
		os.Exit(3)
	}

	// Concat machine ID and username and generate SHA1 hash as beacon_id
	concat_id := machine_id + beacon_username
	base_sha := sha1.New()
	base_sha.Write([]byte(concat_id))
	beacon_id := hex.EncodeToString(base_sha.Sum(nil))
	return beacon_id
}
