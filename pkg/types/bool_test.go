package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool_UnmarshalJSON(t *testing.T) {
	testBool := Bool(true)
	jsonStr := `{"check":true, "check_str":"true"}`
	b := struct {
		BoolField    Bool `json:"check"`
		BoolStrField Bool `json:"check_str"`
	}{}
	err := json.Unmarshal([]byte(jsonStr), &b)
	assert.NoError(t, err)
	assert.Equal(t, testBool, b.BoolField)
	assert.Equal(t, testBool, b.BoolStrField)
}
