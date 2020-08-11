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
func PortsStreamJSONReader(r io.Reader, errors ErrorChannel) PortsChannel {
	portsChan := make(PortsChannel)

	// Reading the JSON files in a streaming fashion.
	go func() {
		defer close(portsChan)

		decoder := json.NewDecoder(r)

		// Reading the first token.
		t, err := decoder.Token()
		if err != nil {
			errors <- err
			return
		}

		// Checking it's an object delimiter '{'.
		dt, ok := t.(json.Delim)
		if !ok || dt != '{' {
			errors <- JSONParseError("Invalid JSON format.")
			return
		}

		for decoder.More() {
			t, err = decoder.Token()
			if err != nil {
				errors <- (err)
			}

			if portID, ok := t.(string); ok && len(portID) == 5 {
				port := &models.Port{}

				if err := decoder.Decode(port); err == nil {
					port.ID = portID
					portsChan <- port
				}
			} else {
				errors <- JSONInvalidPortID(t)
			}
		}
	}()

	return portsChan
}
