package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/x1n13y84issmd42/portobello/PortClient/service"
	"github.com/x1n13y84issmd42/portobello/PortClient/source"
	"github.com/x1n13y84issmd42/portobello/shared/errors"
)

type objectHandler func(*http.Request) (interface{}, uint, error)

// Server is a REST server.
type Server struct {
	Ports service.Ports

	ImportGoing    bool
	ImportProgress uint
}

// New creates a new Server instance.
func New(portsService service.Ports) *Server {
	server := &Server{
		Ports: portsService,
	}
	return server
}

// Serve listens for incoming requests.
func (server *Server) Serve(host string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/ports/", server.JSONHandler(server.HandlePorts))
	mux.HandleFunc("/import", server.JSONHandler(server.HandleImport))

	fmt.Printf("Starting the REST server @ %s...\n", host)

	http.ListenAndServe(fmt.Sprintf("%s", host), mux)
}

// JSONError is a JSON representation of an error message.
type JSONError struct {
	Error string `json:"error"`
}

// JSONHandler creates an http.Handler from an object handler.
func (server *Server) JSONHandler(handler objectHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, status, err := handler(r)

		w.Header().Add("Content-Type", "application/json")

		if err != nil {
			resp = JSONError{Error: err.Error()}
		}

		json, err := json.Marshal(resp)
		if err != nil {
			status = http.StatusInternalServerError
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(int(status))
		w.Write(json)
	}
}

// HandlePorts is a /ports/{portID} route handler.
func (server *Server) HandlePorts(r *http.Request) (interface{}, uint, error) {
	portID := strings.TrimPrefix(r.URL.Path, "/ports/")
	if portID == "" {
		return nil, 404, nil
	}

	fmt.Printf("Port '%s' is requested.\n", portID)

	port, err := server.Ports.GetPort(portID)
	if err != nil {
		fmt.Printf("Error %#v\n", err)
		if errNotFound, ok := err.(errors.ErrPortNotFound); ok {
			return nil, 404, errNotFound
		}

		return nil, 500, err
	}

	return port, 200, nil
}

// ImportProgress is a JSON response for import progress.
type ImportProgress struct {
	Progress uint
}

// HandleImport is a /import route handler.
// It starts (or doesn't if already started) the port import process.
func (server *Server) HandleImport(r *http.Request) (interface{}, uint, error) {
	if !server.ImportGoing {
		fmt.Printf("Importing the ports file.\n")

		progress, err := source.ImportPorts("ports.json", source.PortsStreamJSONReader, server.Ports)
		if err != nil {
			return nil, 500, err
		}

		server.ImportGoing = true
		server.ImportProgress = 0

		go func() {
			for server.ImportProgress = range progress {
				//
			}

			fmt.Printf("Done importing.\n")
			server.ImportGoing = false
		}()

		return "Working!", 202, nil
	}

	return ImportProgress{Progress: server.ImportProgress}, 200, nil
}
