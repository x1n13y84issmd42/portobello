package source_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/portobello/PortClient/source"
	"github.com/x1n13y84issmd42/portobello/shared/models"
)

func Test_PortsStreamJSONReader(T *testing.T) {
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

	T.Run("Alright", func(T *testing.T) {
		r := strings.NewReader(`{
			"AEAJM": {
			  "name": "Ajman",
			  "city": "Ajman",
			  "country": "United Arab Emirates",
			  "alias": [],
			  "regions": [],
			  "coordinates": [
				55.5136433,
				25.4052165
			  ],
			  "province": "Ajman",
			  "timezone": "Asia/Dubai",
			  "unlocs": [
				"AEAJM"
			  ],
			  "code": "52000"
			},
			"AEAUH": {
			  "name": "Abu Dhabi",
			  "coordinates": [
				54.37,
				24.47
			  ],
			  "city": "Abu Dhabi",
			  "province": "Abu Z¸aby [Abu Dhabi]",
			  "country": "United Arab Emirates",
			  "alias": [],
			  "regions": [],
			  "timezone": "Asia/Dubai",
			  "unlocs": [
				"AEAUH"
			  ],
			  "code": "52001"
			}
		}`)

		errorChan := make(source.ErrorChannel)
		portsChan := source.PortsStreamJSONReader(r, errorChan)

		expected := []*models.Port{
			{
				ID:      "AEAJM",
				Name:    "Ajman",
				City:    "Ajman",
				Country: "United Arab Emirates",
				Alias:   []string{},
				Regions: []string{},
				Coordinates: []float64{
					55.5136433,
					25.4052165,
				},
				Province: "Ajman",
				Timezone: "Asia/Dubai",
				Unlocs: []string{
					"AEAJM",
				},
				Code: "52000",
			},
			{
				ID:      "AEAUH",
				Name:    "Abu Dhabi",
				City:    "Abu Dhabi",
				Country: "United Arab Emirates",
				Alias:   []string{},
				Regions: []string{},
				Coordinates: []float64{
					54.37,
					24.47,
				},
				Province: "Abu Z¸aby [Abu Dhabi]",
				Timezone: "Asia/Dubai",
				Unlocs: []string{
					"AEAUH",
				},
				Code: "52001",
			},
		}

		var actual []*models.Port

		for p := range portsChan {
			actual = append(actual, p)
		}

		assert.Equal(T, expected, actual)
	})
}
