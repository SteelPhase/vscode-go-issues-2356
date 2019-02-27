package thing

import "errors"

// Errors
var (
	ErrSimpleError = errors.New("sighs")
)

// Stuff is an example struct
type Stuff struct {
}

// NewStuff is an example constructor
func NewStuff() *Stuff {
	return &Stuff{}
}

// TestThing is an example function
func (s *Stuff) TestThing() error {
	return ErrSimpleError
}
