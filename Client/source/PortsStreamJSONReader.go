package source

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/x1n13y84issmd42/portobello/shared/models"
)

// PortsStreamJSONReader reads a JSON file containing ports data.
func PortsStreamJSONReader(filePath string) (PortsChannel, error) {
	ch := make(PortsChannel)

	file, fileErr := os.Open(filePath)
	if fileErr != nil {
		return nil, fileErr
	}

	go func() {
		fmt.Println("here")
		defer close(ch)
		decoder := json.NewDecoder(file)

		var err error

		// Skipping the '{'
		t, err := decoder.Token()
		if err != nil {
			panic(err)
		}

		for decoder.More() {
			t, err = decoder.Token()

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
