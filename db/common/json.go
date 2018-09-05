package common

import "encoding/json"

type JsonArrayString string

func (o JsonArrayString) Unmarshal() (result []string) {
	json.Unmarshal([]byte(o), result)

	return
}
