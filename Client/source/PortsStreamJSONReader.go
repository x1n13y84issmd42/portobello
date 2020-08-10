package source

import (
	"encoding/json"
	"os"

	"github.com/x1n13y84issmd42/portobello/shared/models"
)

// PortsStreamJSONReader reads a JSON file containing ports data
// and emits parsed ports via the returned channel.
// Working in a streaming fashion, it works with arbitrarily
// large JSON files without running out of memory.
func PortsStreamJSONReader(filePath string) (PortsChannel, error) {
	ch := make(PortsChannel)

	file, fileErr := os.Open(filePath)
	if fileErr != nil {
		return nil, fileErr
	}

	// Reading the JSON files in a streaming fashion.
	go func() {
		defer close(ch)

		decoder := json.NewDecoder(file)

		// Skipping the '{'
		t, err := decoder.Token()
		if err != nil {
			panic(err)
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
			}
		}
	}()

	return ch, nil
}
