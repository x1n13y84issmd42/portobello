package source_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/portobello/shared/models"
	"github.com/x1n13y84issmd42/portobello/shared/proto"
)

func Test_Port(T *testing.T) {
	T.Run("Port->Proto", func(T *testing.T) {
		port := &models.Port{
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
		}

		expected := &proto.Port{
			ID:      "AEAJM",
			Name:    "Ajman",
			City:    "Ajman",
			Country: "United Arab Emirates",
			Alias:   []string{},
			Regions: []string{},
			Coordinates: []float32{
				55.5136433,
				25.4052165,
			},
			Province: "Ajman",
			Timezone: "Asia/Dubai",
			Unlocs: []string{
				"AEAJM",
			},
			Code: "52000",
		}

		assert.Equal(T, expected, port.Proto())
	})

	T.Run("Proto->Port", func(T *testing.T) {
		expected := &models.Port{
			ID:      "AEAJM",
			Name:    "Ajman",
			City:    "Ajman",
			Country: "United Arab Emirates",
			Alias:   []string{},
			Regions: []string{},
			Coordinates: []float64{
				55.51364517211914,
				25.405216217041016,
			},
			Province: "Ajman",
			Timezone: "Asia/Dubai",
			Unlocs: []string{
				"AEAJM",
			},
			Code: "52000",
		}

		protoPort := &proto.Port{
			ID:      "AEAJM",
			Name:    "Ajman",
			City:    "Ajman",
			Country: "United Arab Emirates",
			Alias:   []string{},
			Regions: []string{},
			Coordinates: []float32{
				55.5136433,
				25.4052165,
			},
			Province: "Ajman",
			Timezone: "Asia/Dubai",
			Unlocs: []string{
				"AEAJM",
			},
			Code: "52000",
		}

		assert.Equal(T, expected, models.NewPortFromProto(protoPort))
	})
}
