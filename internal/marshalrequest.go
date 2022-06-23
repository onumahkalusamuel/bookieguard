package internal

import (
	"bytes"
	"encoding/json"

	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/pkg"
)

func MarshalRequest(data config.BodyStructure) *bytes.Buffer {

	data["hashedID"] = config.HashedID

	json_data, _ := json.Marshal(data)
	encrypted := pkg.Encrypt(string(json_data), config.Key)
	d, _ := json.Marshal(config.BodyStructure{"data": string(encrypted)})

	return bytes.NewBuffer(d)
}
