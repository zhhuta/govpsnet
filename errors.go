//errors.go
package govpsnet

import (
	//"errors"
	"fmt"
)

type Errors struct {
	Message string
}

type ErrorApi struct {
	StatusCode int
	Errors     `json:"errors"`
}

func (e *ErrorApi) Error() string {
	return fmt.Sprintf("StatusCode %d, Message: %s", e.StatusCode, e.Errors)
}
