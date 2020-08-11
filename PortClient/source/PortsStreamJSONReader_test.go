package source_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/portobello/PortClient/source"
)

func TestPortsStreamJSONReader(T *testing.T) {
	T.Run("InvalidJSON", func(T *testing.T) {
		r := strings.NewReader("TOTALLY INVALID JSON")

		errorChan := make(source.ErrorChannel)
		source.PortsStreamJSONReader(r, errorChan)

		expected := "invalid character 'T' looking for beginning of value"
		actual := <-errorChan

		assert.Equal(T, expected, actual.Error())
	})

	T.Run("InvalidFormat", func(T *testing.T) {
		r := strings.NewReader("1")

		errorChan := make(source.ErrorChannel)
		source.PortsStreamJSONReader(r, errorChan)

		expected := source.JSONParseError("Invalid JSON format.")
		actual := <-errorChan

		assert.Equal(T, expected, actual)
	})

	T.Run("InvalidFormat", func(T *testing.T) {
		r := strings.NewReader("[1]")

		errorChan := make(source.ErrorChannel)
		source.PortsStreamJSONReader(r, errorChan)

		expected := source.JSONParseError("Invalid JSON format.")
		actual := <-errorChan

		assert.Equal(T, expected, actual)
	})

	T.Run("InvalidPortID", func(T *testing.T) {
		r := strings.NewReader(`{"666":{}}`)

		errorChan := make(source.ErrorChannel)
		source.PortsStreamJSONReader(r, errorChan)

		expected := source.JSONInvalidPortID("666")
		actual := <-errorChan

		assert.Equal(T, expected, actual)
	})
}
