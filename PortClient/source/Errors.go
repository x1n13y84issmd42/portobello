package source

import "fmt"

// ErrJSONParse is raised when there is an error in JSON data.
type ErrJSONParse struct {
	Reason string
}

func (err ErrJSONParse) Error() string {
	return fmt.Sprintf("JSON parse error: %s", err.Reason)
}

// JSONParseError creates a new ErrJSONParse error instance.
func JSONParseError(r string) ErrJSONParse {
	return ErrJSONParse{
		Reason: r,
	}
}

// ErrJSONInvalidPortID is raised when there is invalid
// token encountered when a string port ID is expected.
type ErrJSONInvalidPortID struct {
	Token interface{}
}

func (err ErrJSONInvalidPortID) Error() string {
	return fmt.Sprintf("Encountered invalid token where a port ID is expected. Must be string, but it's %T(%v)", err.Token, err.Token)
}

// JSONInvalidPortID creates a new ErrJSONInvalidPortID error instance.
func JSONInvalidPortID(token interface{}) ErrJSONInvalidPortID {
	return ErrJSONInvalidPortID{
		Token: token,
	}
}
