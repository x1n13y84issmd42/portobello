package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/portobello/PortService/storage"
	"github.com/x1n13y84issmd42/portobello/shared/errors"
	"github.com/x1n13y84issmd42/portobello/shared/models"
)

func Test_MemPorts(T *testing.T) {
	T.Run("GetPort/PortNotFound", func(T *testing.T) {
		ports := storage.NewMemPorts()

		portID := "XXYYZZ"
		expected := errors.PortNotFound(portID, "memory storage")
		port, actual := ports.Get(portID)

		assert.Nil(T, port)
		assert.Equal(T, expected, actual)
	})

	T.Run("Add", func(T *testing.T) {
		ports := storage.NewMemPorts()
		port := &models.Port{
			ID:      "UAODS",
			Name:    "Odessa",
			City:    "Odessa",
			Country: "Ukraine",
			Coordinates: []float64{
				30.723308563232422,
				46.48252487182617,
			},
			Province: "Odessa Oblast",
			Timezone: "Europe/Kiev",
			Unlocs: []string{
				"UAODS",
			},
			Code: "46275",
		}

		err := ports.Add(port)
		assert.Nil(T, err)

		actualPort, err := ports.Get("UAODS")
		assert.Nil(T, err)
		assert.Equal(T, port, actualPort)
	})
}
