package source_test

import (
	goerrors "errors"
	"fmt"
	"io"
	"sort"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/portobello/PortClient/source"
	"github.com/x1n13y84issmd42/portobello/shared/models"
)

func Mock_PortsReader(r io.Reader, errors source.ErrorChannel) source.PortsChannel {
	ch := make(source.PortsChannel)

	go func() {
		defer close(ch)

		ch <- &models.Port{}
		ch <- &models.Port{}
		ch <- &models.Port{}
		ch <- &models.Port{}

	}()

	return ch
}

func Mock_PortsReader_Error(r io.Reader, errors source.ErrorChannel) source.PortsChannel {
	ch := make(source.PortsChannel)

	go func() {
		defer close(ch)

		ch <- &models.Port{}
		errors <- goerrors.New("err 1")
		ch <- &models.Port{}
		errors <- goerrors.New("err 2")
		ch <- &models.Port{}
		errors <- goerrors.New("err 3")
		ch <- &models.Port{}

	}()

	return ch
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

	T.Run("Errors", func(T *testing.T) {
		expected := []string{
			"err 1",
			"err 2",
			"err 3",
			"oops_1",
			"oops_2",
			"oops_3",
			"oops_4",
		}

		ports := &Mock_PortsService_Counter{
			DoError: true,
		}

		progressChan, errorChan, _ := source.ImportPorts(strings.NewReader(""), Mock_PortsReader_Error, ports)

		wg := sync.WaitGroup{}
		wg.Add(2)

		var progress uint = 0
		go func() {
			for progress = range progressChan {
				//
			}

			wg.Done()
		}()

		var actual []string

		go func() {
			for err := range errorChan {
				actual = append(actual, err.Error())
			}

			wg.Done()
		}()

		wg.Wait()

		sort.Strings(actual)

		assert.Equal(T, progress, uint(0))
		assert.Equal(T, expected, actual)
	})

}
