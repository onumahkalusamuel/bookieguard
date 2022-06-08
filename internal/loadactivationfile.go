package internal

import (
	"encoding/json"
	"os"

	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/pkg"
)

func LoadActivationFile() (bool, config.BodyStructure) {

	f, err := os.ReadFile(config.ActivationFile)
	if err != nil || len(string(f)) < 36 {
		return false, config.BodyStructure{}
	}

	decrypted := pkg.Decrypt(string(f), config.Key)

	var unmarshalled config.BodyStructure

	json.Unmarshal([]byte(decrypted), &unmarshalled)

	return true, unmarshalled

}
