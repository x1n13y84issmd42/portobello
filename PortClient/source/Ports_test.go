package source_test

import (
	goerrors "errors"
	"fmt"
	"io"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/portobello/PortClient/source"
	"github.com/x1n13y84issmd42/portobello/shared/models"
)

func Mock_PortsReader(r io.Reader) (source.PortsChannel, error) {
	ch := make(source.PortsChannel)

	go func() {
		defer close(ch)

		ch <- &models.Port{}
		ch <- &models.Port{}
		ch <- &models.Port{}
		ch <- &models.Port{}

	}()

	return ch, nil
}

func Mock_PortsReader_Error(r io.Reader) (source.PortsChannel, error) {
	return nil, goerrors.New("oops")
}

type Mock_PortsService_Counter struct {
	AddedCounter uint
	ErrorCounter uint
	DoError      bool
}

// Close ...
func (ports *Mock_PortsService_Counter) Close() {
	//
}

// AddPort ...
func (ports *Mock_PortsService_Counter) AddPort(port *models.Port) error {
	if ports.DoError {
		ports.ErrorCounter++
		return fmt.Errorf("oops_%d", ports.ErrorCounter)
	}

	ports.AddedCounter++
	return nil
}

// GetPort ...
func (ports *Mock_PortsService_Counter) GetPort(id models.PortID) (*models.Port, error) {
	return nil, nil
}

func Test_ImportPorts(T *testing.T) {
	T.Run("OK", func(T *testing.T) {
		ports := &Mock_PortsService_Counter{}
		progressChan, _, err := source.ImportPorts(strings.NewReader(""), Mock_PortsReader, ports)

		var progress uint = 0

		for progress = range progressChan {
			//
		}

		assert.Nil(T, err)
		assert.Equal(T, uint(4), progress)
		assert.Equal(T, progress, ports.AddedCounter)
	})

	T.Run("ErrorChan", func(T *testing.T) {
		expected := []error{
			goerrors.New("oops_1"),
			goerrors.New("oops_2"),
			goerrors.New("oops_3"),
			goerrors.New("oops_4"),
		}

		ports := &Mock_PortsService_Counter{
			DoError: true,
		}

		progressChan, errorChan, _ := source.ImportPorts(strings.NewReader(""), Mock_PortsReader, ports)

		wg := sync.WaitGroup{}
		wg.Add(1)

		var progress uint = 0
		go func() {
			for progress = range progressChan {
				//
			}

			wg.Done()
		}()

		var actual []error
		go func() {
			for err := range errorChan {
				actual = append(actual, err)
			}

			wg.Done()
		}()

		wg.Wait()

		assert.Equal(T, progress, uint(0))
		assert.Equal(T, expected, actual)
	})

	T.Run("Error", func(T *testing.T) {
		ports := &Mock_PortsService_Counter{}

		expected := goerrors.New("oops")

		progressChan, errorChan, actual := source.ImportPorts(strings.NewReader(""), Mock_PortsReader_Error, ports)

		assert.Nil(T, progressChan)
		assert.Nil(T, errorChan)
		assert.Equal(T, expected, actual)
	})

}
