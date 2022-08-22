package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Bool bool

func (b Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(b))
}

func (b *Bool) UnmarshalJSON(data []byte) error {
	asString := strings.ToLower(strings.Trim(string(data), `"`))
	switch asString {
	case "true":
		*b = true
	case "false":
		*b = false
	default:
		return errors.New(fmt.Sprintf("Boolean unmarshal error: invalid input %s", asString))
	}
	return nil
}
