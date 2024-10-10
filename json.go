package snapshot

import (
	"encoding/json"
	"io"
	"testing"
)

func MatchJSON(t *testing.T, value any) {
	t.Helper()
	Match(t, jsonRenderer{}, value)
}

type jsonRenderer struct{}

func (jsonRenderer) Render(w io.Writer, v interface{}) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	return encoder.Encode(v)
}

func (jsonRenderer) Ext() string {
	return ".json"
}
