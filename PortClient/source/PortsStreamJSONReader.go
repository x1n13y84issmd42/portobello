package source

import (
	"encoding/json"
	"io"

	"github.com/x1n13y84issmd42/portobello/shared/models"
)

// PortsStreamJSONReader reads a JSON file containing ports data
// and emits parsed ports via the returned channel.
// Working in a streaming fashion, it works with arbitrarily
// large JSON files without running out of memory.
func PortsStreamJSONReader(r io.Reader) (PortsChannel, error) {
	ch := make(PortsChannel)

	// Reading the JSON files in a streaming fashion.
	go func() {
		defer close(ch)

		decoder := json.NewDecoder(r)

		// Skipping the '{'
		t, err := decoder.Token()
		if err != nil {
			panic(err)
		}

		if st, ok := t.(string); ok && st != "{" {
			panic(JSONParseError("Invalid JSON format."))
		}

		for decoder.More() {
			t, err = decoder.Token()
			if err != nil {
				panic(err)
			}

			if portID, ok := t.(string); ok {
				port := &models.Port{}

				if err := decoder.Decode(port); err == nil {
					port.ID = portID
					ch <- port
				}
			} else {
				panic(JSONInvalidPortID(t))
			}
			//TODO: else {panic with wrong key type error}
		}
	}()

	return ch, nil
}
