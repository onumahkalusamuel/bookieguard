package internal

import (
	"encoding/json"
	"os"

	"github.com/onumahkalusamuel/bookieguard/config"
)

func MockServerResponse(config.BodyStructure) config.BodyStructure {

	var holder config.BodyStructure

	res, _ := os.ReadFile(config.MockData)

	json.Unmarshal(res, &holder)

	holder["email"] = config.Email
	holder["shop"] = config.Shop
	holder["computerName"] = config.ComputerName
	holder["hashedID"] = config.HashedID

	return holder
}
