package internal

import (
	"encoding/json"
	"net/http"

	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/pkg"
)

func UnmarshalResponse(resp *http.Response) (map[string]string, error) {

	var res map[string]string

	err := json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return map[string]string{}, err
	}

	decrypted := pkg.Decrypt(res["data"], config.Key)

	var newMap map[string]string

	json.Unmarshal([]byte(decrypted), &newMap)

	return newMap, nil
}
